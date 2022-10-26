// Copyright (c) 2022 Wibowo Arindrarto <contact@arindrarto.dev>
// SPDX-License-Identifier: BSD-3-Clause

package store

import (
	"context"
	"database/sql"
)

// EditFeed updates fields of an feed.
func (s *SQLite) EditFeeds(
	ctx context.Context,
	ops []*FeedEditOp,
) ([]*Feed, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	updateFunc := func(ctx context.Context, tx *sql.Tx, op *FeedEditOp) (*Feed, error) {
		if err := setFeedTitle(ctx, tx, op.DBID, op.Title); err != nil {
			return nil, err
		}
		if err := setFeedDescription(ctx, tx, op.DBID, op.Description); err != nil {
			return nil, err
		}
		if err := setFeedTags(ctx, tx, op.DBID, op.Tags); err != nil {
			return nil, err
		}
		if err := setFeedIsStarred(ctx, tx, op.DBID, op.IsStarred); err != nil {
			return nil, err
		}
		return getFeed(ctx, tx, op.DBID)
	}

	var entries = make([]*Feed, len(ops))
	dbFunc := func(ctx context.Context, tx *sql.Tx) error {
		for i, op := range ops {
			feed, err := updateFunc(ctx, tx, op)
			if err != nil {
				return err
			}
			entries[i] = feed
		}
		return nil
	}

	fail := failF("SQLite.EditFeed")

	err := s.withTx(ctx, dbFunc)
	if err != nil {
		return nil, fail(err)
	}
	return entries, nil
}

func getFeed(ctx context.Context, tx *sql.Tx, feedDBID DBID) (*Feed, error) {

	sql1 := `
		SELECT
			f.id AS id,
			f.title AS title,
			f.description AS description,
			f.feed_url AS feed_url,
			f.site_url AS site_url,
			f.is_starred AS is_starred,
			f.subscription_time AS subscription_time,
			f.update_time AS update_time,
			json_group_array(fc.name) FILTER (WHERE fc.name IS NOT NULL) AS tags
		FROM
			feeds f
			LEFT JOIN feeds_x_feed_tags fxfc ON fxfc.feed_id = f.id
			LEFT JOIN feed_tags fc ON fxfc.feed_tag_id = fc.id
		WHERE
			f.id = ?
		GROUP BY
			f.id
		ORDER BY
			COALESCE(f.update_time, f.subscription_time) DESC
`
	scanRow := func(row *sql.Row) (*Feed, error) {
		var feed Feed
		if err := row.Scan(
			&feed.DBID,
			&feed.Title,
			&feed.Description,
			&feed.FeedURL,
			&feed.SiteURL,
			&feed.IsStarred,
			&feed.Subscribed,
			&feed.Updated,
			&feed.Tags,
		); err != nil {
			return nil, err
		}
		if len(feed.Tags) == 0 {
			feed.Tags = nil
		}
		return &feed, nil
	}

	stmt1, err := tx.PrepareContext(ctx, sql1)
	if err != nil {
		return nil, err
	}
	defer stmt1.Close()

	return scanRow(stmt1.QueryRowContext(ctx, feedDBID))
}

var (
	setFeedTitle       = setTableField[string]("feeds", "title")
	setFeedDescription = setTableField[string]("feeds", "description")
	setFeedIsStarred   = setTableField[bool]("feeds", "is_starred")
	setFeedSiteURL     = setTableField[string]("feeds", "site_url")
)

func setFeedTags(
	ctx context.Context,
	tx *sql.Tx,
	feedDBID DBID,
	tags *[]string,
) error {

	if tags == nil {
		return nil
	}

	sql1 := `DELETE FROM feeds_x_feed_tags WHERE feed_id = ?`
	stmt1, err := tx.PrepareContext(ctx, sql1)
	if err != nil {
		return err
	}
	defer stmt1.Close()

	if _, err = stmt1.ExecContext(ctx); err != nil {
		return err
	}

	if err = addFeedTags(ctx, tx, feedDBID, *tags); err != nil {
		return err
	}

	sql2 := `
		DELETE
			feed_tags
		WHERE
			id IN (
				SELECT
					fc.id
				FROM
					feed_tags fc
					LEFT JOIN feeds_x_feed_tags fxfc ON fxfc.feed_tag_id = fc.id
				WHERE
					fxfc.feed_id IS NULL
			)
	`
	stmt2, err := tx.PrepareContext(ctx, sql2)
	if err != nil {
		return err
	}
	defer stmt2.Close()

	return nil
}
