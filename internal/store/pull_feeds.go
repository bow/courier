package store

import (
	"context"
	"database/sql"
	"sync"

	"github.com/rs/zerolog/log"
)

func (s *SQLite) PullFeeds(ctx context.Context) <-chan PullResult {

	fail := failF("SQLite.PullFeeds")
	c := make(chan PullResult)

	go func() {
		defer close(c)

		tx, err := s.db.BeginTx(ctx, nil)
		if err != nil {
			c <- newPullResultFromErr(fail(err))
			return
		}

		rb := func(tx *sql.Tx) {
			if rerr := tx.Rollback(); rerr != nil {
				log.Error().Err(rerr).Msg("failed to roll back transaction")
			}
		}

		defer func() {
			if p := recover(); p != nil {
				rb(tx)
				panic(p)
			}
			if err != nil {
				rb(tx)
			} else {
				if err = tx.Commit(); err != nil {
					c <- newPullResultFromErr(fail(err))
				}
			}
		}()

		pks, err := getAllPullKeys(ctx, tx)
		if err != nil {
			c <- newPullResultFromErr(fail(err))
			return
		}
		if len(pks) == 0 {
			c <- PullResult{status: pullSuccess}
			return
		}

		chs := make([]<-chan PullResult, len(pks))
		for i, pk := range pks {
			chs[i] = pullNewFeedEntries(ctx, tx, pk, s.parser)
		}

		for pr := range merge(chs) {
			pr := pr
			if pr.Error() != nil {
				pr.err = fail(pr.err)
			}
			c <- pr
		}
	}()

	return c
}

// PullResult is a container for a pull operation.
type PullResult struct {
	pk     pullKey
	status pullStatus
	ok     *Feed
	err    error
}

func (msg PullResult) Result() *Feed {
	if msg.status == pullSuccess {
		return msg.ok
	}
	return nil
}

func (msg PullResult) Error() error {
	if msg.status == pullFail {
		return msg.err
	}
	return nil
}

func newPullResultFromErr(err error) PullResult {
	return PullResult{status: pullFail, err: err}
}

type pullStatus int

const (
	pullSuccess pullStatus = iota
	pullFail
)

type pullKey struct {
	feedDBID DBID
	feedURL  string
}

func (pk pullKey) ok(feed *Feed) PullResult {
	return PullResult{pk: pk, status: pullSuccess, ok: feed, err: nil}
}

func (pk pullKey) err(e error) PullResult {
	return PullResult{pk: pk, status: pullFail, ok: nil, err: e}
}

var setFeedUpdateTime = setTableField[string]("feeds", "update_time")

func getAllPullKeys(ctx context.Context, tx *sql.Tx) ([]pullKey, error) {

	sql1 := `SELECT id, feed_url FROM feeds`

	scanRow := func(rows *sql.Rows) (pullKey, error) {
		var pk pullKey
		err := rows.Scan(&pk.feedDBID, &pk.feedURL)
		return pk, err
	}

	stmt1, err := tx.PrepareContext(ctx, sql1)
	if err != nil {
		return nil, err
	}

	rows, err := stmt1.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	pks := make([]pullKey, 0)
	for rows.Next() {
		pk, err := scanRow(rows)
		if err != nil {
			return nil, err
		}
		pks = append(pks, pk)
	}

	return pks, nil
}

func pullNewFeedEntries(
	ctx context.Context,
	tx *sql.Tx,
	pk pullKey,
	parser FeedParser,
) chan PullResult {

	pullf := func() PullResult {

		gfeed, err := parser.ParseURLWithContext(pk.feedURL, ctx)
		if err != nil {
			return pk.err(err)
		}

		updateTime := serializeTime(resolveFeedUpdateTime(gfeed))
		if err = setFeedUpdateTime(ctx, tx, pk.feedDBID, updateTime); err != nil {
			return pk.err(err)
		}

		if len(gfeed.Items) == 0 {
			return pk.ok(nil)
		}

		if err = upsertEntries(ctx, tx, pk.feedDBID, gfeed.Items); err != nil {
			return pk.err(err)
		}

		entries, err := getAllFeedEntries(ctx, tx, pk.feedDBID, pointer(false))
		if err != nil {
			return pk.err(err)
		}

		feed, err := getFeed(ctx, tx, pk.feedDBID)
		if err != nil {
			return pk.err(err)
		}

		feed.Entries = entries

		return pk.ok(feed)
	}

	ic := make(chan PullResult)
	go func() {
		defer close(ic)
		ic <- pullf()
	}()

	oc := make(chan PullResult)
	go func() {
		defer close(oc)
		select {
		case <-ctx.Done():
			oc <- pk.err(ctx.Err())
		case msg := <-ic:
			oc <- msg
		}
	}()

	return oc
}

func merge[T any](chs []<-chan T) chan T {
	var (
		wg     sync.WaitGroup
		merged = make(chan T, len(chs))
	)

	forward := func(ch <-chan T) {
		for msg := range ch {
			merged <- msg
		}
		wg.Done()
	}

	wg.Add(len(chs))
	for _, ch := range chs {
		go forward(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}
