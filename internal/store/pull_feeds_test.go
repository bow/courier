// Copyright (c) 2022 Wibowo Arindrarto <contact@arindrarto.dev>
// SPDX-License-Identifier: BSD-3-Clause

package store

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/bow/iris/internal"
	gomock "github.com/golang/mock/gomock"
	"github.com/mmcdole/gofeed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPullFeedsAllOkEmptyDB(t *testing.T) {
	t.Parallel()

	a := assert.New(t)
	r := require.New(t)
	st := newTestStore(t)

	r.Equal(0, st.countFeeds())

	st.parser.EXPECT().
		ParseURLWithContext(gomock.Any(), gomock.Any()).
		MaxTimes(0)

	c := st.PullFeeds(context.Background(), nil)
	a.Empty(c)
}

func TestPullFeedsAllOkEmptyEntries(t *testing.T) {
	t.Parallel()

	a := assert.New(t)
	r := require.New(t)
	st := newTestStore(t)

	dbFeeds := []*feedRecord{
		{
			title:   "Feed A",
			feedURL: "http://a.com/feed.xml",
			updated: toNullTime(mustTime(t, "2022-03-19T16:23:18.600+02:00")),
			entries: []*entryRecord{},
		},
		{
			title:   "Feed X",
			feedURL: "http://x.com/feed.xml",
			updated: toNullTime(mustTime(t, "2022-04-20T16:32:30.760+02:00")),
			entries: []*entryRecord{},
		},
	}

	st.addFeeds(dbFeeds)
	r.Equal(2, st.countFeeds())

	st.parser.EXPECT().
		ParseURLWithContext(dbFeeds[0].feedURL, gomock.Any()).
		MaxTimes(1).
		Return(toGFeed(t, dbFeeds[0]), nil)

	st.parser.EXPECT().
		ParseURLWithContext(dbFeeds[1].feedURL, gomock.Any()).
		MaxTimes(1).
		Return(toGFeed(t, dbFeeds[1]), nil)

	c := st.PullFeeds(context.Background(), nil)

	got := make([]internal.PullResult, 0)
	for res := range c {
		got = append(got, res)
	}

	want := []internal.PullResult{
		internal.NewPullResultFromFeed(
			&dbFeeds[0].feedURL,
			nil,
		),
		internal.NewPullResultFromFeed(
			&dbFeeds[1].feedURL,
			nil,
		),
	}

	a.ElementsMatch(want, got)
}

func TestPullFeedsAllOkNoNewEntries(t *testing.T) {
	t.Parallel()

	a := assert.New(t)
	r := require.New(t)
	st := newTestStore(t)

	dbFeeds := []*feedRecord{
		{
			title:   "Feed A",
			feedURL: "http://a.com/feed.xml",
			updated: toNullTime(mustTime(t, "2022-03-19T16:23:18.600+02:00")),
			entries: []*entryRecord{
				{
					title:   "Entry A1",
					extID:   "A1",
					isRead:  true,
					updated: toNullString("2022-07-16T23:39:07.383+02:00"),
					url:     toNullString("http://a.com/a1.html"),
				},
				{
					title:   "Entry A2",
					extID:   "A2",
					isRead:  true,
					updated: toNullString("2022-07-16T23:42:24.988+02:00"),
					url:     toNullString("http://a.com/a2.html"),
				},
			},
		},
		{
			title:   "Feed X",
			feedURL: "http://x.com/feed.xml",
			updated: toNullTime(mustTime(t, "2022-04-20T16:32:30.760+02:00")),
			entries: []*entryRecord{
				{
					title:   "Entry X1",
					extID:   "X1",
					isRead:  true,
					updated: toNullString("2022-07-16T23:43:12.759+02:00"),
					url:     toNullString("http://x.com/x1.html"),
				},
			},
		},
	}

	st.addFeeds(dbFeeds)
	r.Equal(2, st.countFeeds())

	pulledFeeds := []*feedRecord{
		{
			title:   dbFeeds[0].title,
			feedURL: dbFeeds[0].feedURL,
			updated: dbFeeds[0].updated,
			entries: []*entryRecord{
				{
					title:   dbFeeds[0].entries[0].title,
					extID:   dbFeeds[0].entries[0].extID,
					updated: dbFeeds[0].entries[0].updated,
					url:     dbFeeds[0].entries[0].url,
				},
				{
					title:   dbFeeds[0].entries[1].title,
					extID:   dbFeeds[0].entries[1].extID,
					updated: dbFeeds[0].entries[1].updated,
					url:     dbFeeds[0].entries[1].url,
				},
			},
		},
		{
			title:   dbFeeds[1].title,
			feedURL: dbFeeds[1].feedURL,
			updated: dbFeeds[1].updated,
			entries: []*entryRecord{
				{
					title:   dbFeeds[1].entries[0].title,
					extID:   dbFeeds[1].entries[0].extID,
					updated: dbFeeds[1].entries[0].updated,
					url:     dbFeeds[1].entries[0].url,
				},
			},
		},
	}

	st.parser.EXPECT().
		ParseURLWithContext(dbFeeds[0].feedURL, gomock.Any()).
		MaxTimes(1).
		Return(toGFeed(t, pulledFeeds[0]), nil)

	st.parser.EXPECT().
		ParseURLWithContext(dbFeeds[1].feedURL, gomock.Any()).
		MaxTimes(1).
		Return(toGFeed(t, pulledFeeds[1]), nil)

	c := st.PullFeeds(context.Background(), nil)

	got := make([]internal.PullResult, 0)
	for res := range c {
		got = append(got, res)
	}

	want := []internal.PullResult{
		internal.NewPullResultFromFeed(
			&pulledFeeds[0].feedURL,
			nil,
		),
		internal.NewPullResultFromFeed(
			&pulledFeeds[1].feedURL,
			nil,
		),
	}

	a.ElementsMatch(want, got)
}

func TestPullFeedsAllOkSomeNewEntries(t *testing.T) {
	t.Parallel()

	a := assert.New(t)
	r := require.New(t)
	st := newTestStore(t)

	dbFeeds := []*feedRecord{ // nolint:dupl
		{
			title:      "Feed A",
			feedURL:    "http://a.com/feed.xml",
			subscribed: mustTime(t, "2022-07-18T22:04:37Z"),
			lastPulled: mustTime(t, "2022-07-18T22:04:37Z"),
			updated:    toNullTime(mustTime(t, "2022-03-19T16:23:18.600+02:00")),
			entries: []*entryRecord{
				{
					// This entry should not be returned later; 'updated' remains the same.
					title:   "Entry A1",
					extID:   "A1",
					isRead:  true,
					updated: toNullString("2022-07-16T23:39:07.383+02:00"),
					url:     toNullString("http://a.com/a1.html"),
				},
				{
					// This entry should not be returned later; 'updated' remains the same.
					title:   "Entry A2",
					extID:   "A2",
					isRead:  false,
					updated: toNullString("2022-07-16T23:42:24.988+02:00"),
					url:     toNullString("http://a.com/a2.html"),
				},
				{
					// This entry should be returned later; 'updated' will be changed.
					title:   "Entry A3",
					extID:   "A3",
					isRead:  true,
					updated: toNullString("2022-03-18T22:51:49.404+02:00"),
					url:     toNullString("http://a.com/a3.html"),
				},
			},
		},
		{
			title:      "Feed X",
			feedURL:    "http://x.com/feed.xml",
			subscribed: mustTime(t, "2022-07-18T22:04:45Z"),
			lastPulled: mustTime(t, "2022-07-18T22:04:45Z"),
			updated:    toNullTime(mustTime(t, "2022-04-20T16:32:30.760+02:00")),
			entries: []*entryRecord{
				{
					// This entry should not be returned later; 'updated' remains the same.
					title:   "Entry X1",
					extID:   "X1",
					isRead:  true,
					updated: toNullString("2022-07-16T23:43:12.759+02:00"),
					url:     toNullString("http://x.com/x1.html"),
				},
			},
		},
	}

	keys := st.addFeeds(dbFeeds)
	r.Equal(2, st.countFeeds())

	pulledFeeds := []*feedRecord{
		{
			title:   dbFeeds[0].title,
			feedURL: dbFeeds[0].feedURL,
			updated: toNullTime(mustTime(t, "2022-07-18T22:51:49.404+02:00")),
			entries: []*entryRecord{
				{
					title:   dbFeeds[0].entries[0].title,
					extID:   dbFeeds[0].entries[0].extID,
					updated: dbFeeds[0].entries[0].updated,
					url:     dbFeeds[0].entries[0].url,
				},
				{
					title:   dbFeeds[0].entries[1].title,
					extID:   dbFeeds[0].entries[1].extID,
					updated: dbFeeds[0].entries[1].updated,
					url:     dbFeeds[0].entries[1].url,
				},
				{
					title:   dbFeeds[0].entries[2].title,
					extID:   dbFeeds[0].entries[2].extID,
					updated: toNullString("2022-07-19T16:23:18.600+02:00"),
					url:     dbFeeds[0].entries[2].url,
				},
			},
		},
		{
			title:   dbFeeds[1].title,
			feedURL: dbFeeds[1].feedURL,
			updated: toNullTime(mustTime(t, "2022-07-18T22:21:41.647+02:00")),
			entries: []*entryRecord{
				{
					title:   dbFeeds[1].entries[0].title,
					extID:   dbFeeds[1].entries[0].extID,
					updated: dbFeeds[1].entries[0].updated,
					url:     dbFeeds[1].entries[0].url,
				},
				{
					title:   "Entry X2",
					extID:   "X2",
					updated: toNullString("2022-07-18T22:21:41.647+02:00"),
					url:     toNullString("http://x.com/x2.html"),
				},
			},
		},
	}

	st.parser.EXPECT().
		ParseURLWithContext(dbFeeds[0].feedURL, gomock.Any()).
		MaxTimes(1).
		Return(toGFeed(t, pulledFeeds[0]), nil)

	st.parser.EXPECT().
		ParseURLWithContext(dbFeeds[1].feedURL, gomock.Any()).
		MaxTimes(1).
		Return(toGFeed(t, pulledFeeds[1]), nil)

	c := st.PullFeeds(context.Background(), nil)

	got := make([]internal.PullResult, 0)
	for res := range c {
		got = append(got, res)
	}

	feedURL0 := pulledFeeds[0].feedURL
	feedURL1 := pulledFeeds[1].feedURL

	want := []internal.PullResult{
		internal.NewPullResultFromFeed(
			&dbFeeds[0].feedURL,
			&internal.Feed{
				ID:         keys[pulledFeeds[0].title].ID,
				Title:      pulledFeeds[0].title,
				FeedURL:    pulledFeeds[0].feedURL,
				Updated:    st.getFeedUpdateTime(feedURL0),
				Subscribed: st.getFeedSubTime(feedURL0),
				LastPulled: time.Time{},
				Entries: []*internal.Entry{
					{
						ID:        st.getEntryID(feedURL0, pulledFeeds[0].entries[1].extID),
						FeedID:    keys[pulledFeeds[0].title].ID,
						Title:     pulledFeeds[0].entries[1].title,
						ExtID:     pulledFeeds[0].entries[1].extID,
						Updated:   st.getEntryUpdateTime(feedURL0, pulledFeeds[0].entries[1].extID),
						Published: st.getEntryPubTime(feedURL0, pulledFeeds[0].entries[1].extID),
						URL:       fromNullString(pulledFeeds[0].entries[1].url),
						IsRead:    false,
					},
					{
						ID:        st.getEntryID(feedURL0, pulledFeeds[0].entries[2].extID),
						FeedID:    keys[pulledFeeds[0].title].ID,
						Title:     pulledFeeds[0].entries[2].title,
						ExtID:     pulledFeeds[0].entries[2].extID,
						Updated:   st.getEntryUpdateTime(feedURL0, pulledFeeds[0].entries[2].extID),
						Published: st.getEntryPubTime(feedURL0, pulledFeeds[0].entries[2].extID),
						URL:       fromNullString(pulledFeeds[0].entries[2].url),
						IsRead:    false,
					},
				},
			},
		),
		internal.NewPullResultFromFeed(
			&dbFeeds[1].feedURL,
			&internal.Feed{
				ID:         keys[pulledFeeds[1].title].ID,
				Title:      pulledFeeds[1].title,
				FeedURL:    pulledFeeds[1].feedURL,
				Updated:    st.getFeedUpdateTime(feedURL1),
				Subscribed: st.getFeedSubTime(feedURL1),
				LastPulled: time.Time{},
				Entries: []*internal.Entry{
					{
						ID:        st.getEntryID(feedURL1, pulledFeeds[1].entries[1].extID),
						FeedID:    keys[pulledFeeds[1].title].ID,
						Title:     pulledFeeds[1].entries[1].title,
						ExtID:     pulledFeeds[1].entries[1].extID,
						Updated:   st.getEntryUpdateTime(feedURL1, pulledFeeds[1].entries[1].extID),
						Published: st.getEntryPubTime(feedURL1, pulledFeeds[1].entries[1].extID),
						URL:       fromNullString(pulledFeeds[1].entries[1].url),
						IsRead:    false,
					},
				},
			},
		),
	}

	// Sort inner entries first, since ElementsMatch cares about inner array elements order.
	sortPullResultEntries(want)
	sortPullResultEntries(got)

	// Set LastPulled fields to the zero value as this value is always updated on every pull.
	for _, item := range got {
		item.Feed().LastPulled = time.Time{}
	}

	a.ElementsMatch(want, got)
}

func TestPullFeedsSelectedOkSomeNewEntries(t *testing.T) {
	t.Parallel()

	a := assert.New(t)
	r := require.New(t)
	st := newTestStore(t)

	dbFeeds := []*feedRecord{ // nolint:dupl
		// This feed should not be returned later, it is not selected.
		{
			title:      "Feed A",
			feedURL:    "http://a.com/feed.xml",
			subscribed: mustTime(t, "2022-07-18T22:04:37Z"),
			lastPulled: mustTime(t, "2022-07-18T22:04:37Z"),
			updated:    toNullTime(mustTime(t, "2022-03-19T16:23:18.600+02:00")),
			entries: []*entryRecord{
				{
					title:   "Entry A1",
					extID:   "A1",
					isRead:  true,
					updated: toNullString("2022-07-16T23:39:07.383+02:00"),
					url:     toNullString("http://a.com/a1.html"),
				},
				{
					title:   "Entry A2",
					extID:   "A2",
					isRead:  false,
					updated: toNullString("2022-07-16T23:42:24.988+02:00"),
					url:     toNullString("http://a.com/a2.html"),
				},
				{
					title:   "Entry A3",
					extID:   "A3",
					isRead:  true,
					updated: toNullString("2022-03-18T22:51:49.404+02:00"),
					url:     toNullString("http://a.com/a3.html"),
				},
			},
		},
		// This feed should be returned later, it is selected.
		{
			title:      "Feed X",
			feedURL:    "http://x.com/feed.xml",
			subscribed: mustTime(t, "2022-07-18T22:04:45Z"),
			lastPulled: mustTime(t, "2022-07-18T22:04:45Z"),
			updated:    toNullTime(mustTime(t, "2022-04-20T16:32:30.760+02:00")),
			entries: []*entryRecord{
				{
					// This entry should not be returned later; 'updated' remains the same.
					title:   "Entry X1",
					extID:   "X1",
					isRead:  true,
					updated: toNullString("2022-07-16T23:43:12.759+02:00"),
					url:     toNullString("http://x.com/x1.html"),
				},
			},
		},
	}

	keys := st.addFeeds(dbFeeds)
	r.Equal(2, st.countFeeds())

	pulledFeed := &feedRecord{
		title:   dbFeeds[1].title,
		feedURL: dbFeeds[1].feedURL,
		updated: toNullTime(mustTime(t, "2022-07-18T22:21:41.647+02:00")),
		entries: []*entryRecord{
			{
				title:   dbFeeds[1].entries[0].title,
				extID:   dbFeeds[1].entries[0].extID,
				updated: dbFeeds[1].entries[0].updated,
				url:     dbFeeds[1].entries[0].url,
			},
			{
				title:   "Entry X2",
				extID:   "X2",
				updated: toNullString("2022-07-18T22:21:41.647+02:00"),
				url:     toNullString("http://x.com/x2.html"),
			},
		},
	}

	st.parser.EXPECT().
		ParseURLWithContext(dbFeeds[1].feedURL, gomock.Any()).
		MaxTimes(1).
		Return(toGFeed(t, pulledFeed), nil)

	c := st.PullFeeds(context.Background(), []ID{keys[pulledFeed.title].ID})

	got := make([]internal.PullResult, 0)
	for res := range c {
		got = append(got, res)
	}

	want := []internal.PullResult{
		internal.NewPullResultFromFeed(
			&dbFeeds[1].feedURL,
			&internal.Feed{
				ID:         keys[pulledFeed.title].ID,
				Title:      pulledFeed.title,
				FeedURL:    pulledFeed.feedURL,
				Updated:    st.getFeedUpdateTime(pulledFeed.feedURL),
				Subscribed: st.getFeedSubTime(pulledFeed.feedURL),
				LastPulled: time.Time{},
				Entries: []*internal.Entry{
					{
						ID:        st.getEntryID(pulledFeed.feedURL, pulledFeed.entries[1].extID),
						FeedID:    keys[pulledFeed.title].ID,
						Title:     pulledFeed.entries[1].title,
						ExtID:     pulledFeed.entries[1].extID,
						Updated:   st.getEntryUpdateTime(pulledFeed.feedURL, pulledFeed.entries[1].extID),
						Published: st.getEntryPubTime(pulledFeed.feedURL, pulledFeed.entries[1].extID),
						URL:       fromNullString(pulledFeed.entries[1].url),
						IsRead:    false,
					},
				},
			},
		),
	}

	// Sort inner entries first, since ElementsMatch cares about inner array elements order.
	sortPullResultEntries(want)
	sortPullResultEntries(got)

	// Set LastPulled fields to the zero value as this value is always updated on every pull.
	for _, item := range got {
		item.Feed().LastPulled = time.Time{}
	}

	a.ElementsMatch(want, got)
}

func sortPullResultEntries(arr []internal.PullResult) {
	for _, item := range arr {
		sort.SliceStable(
			item.Feed().Entries,
			func(i, j int) bool {
				return item.Feed().Entries[i].ExtID < item.Feed().Entries[j].ExtID
			},
		)
	}
}

func toGFeed(t *testing.T, feed *feedRecord) *gofeed.Feed {
	t.Helper()
	gfeed := gofeed.Feed{
		Title:    feed.title,
		FeedLink: feed.feedURL,
	}
	if !feed.updated.Time.IsZero() {
		gfeed.Updated = feed.updated.Time.UTC().Format(time.RFC3339Nano)
		gfeed.UpdatedParsed = &feed.updated.Time
	}
	for _, entry := range feed.entries {
		item := gofeed.Item{
			GUID:    entry.extID,
			Link:    entry.url.String,
			Title:   entry.title,
			Content: entry.content.String,
		}
		if entry.published.String != "" {
			item.Published = entry.published.String
			item.PublishedParsed = mustTimeP(t, entry.published.String)
		}
		if entry.updated.String != "" {
			item.Updated = entry.updated.String
			item.UpdatedParsed = mustTimeP(t, entry.updated.String)
		}
		gfeed.Items = append(gfeed.Items, &item)
	}
	return &gfeed
}
