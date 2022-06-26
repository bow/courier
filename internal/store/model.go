package store

import (
	"database/sql"
	"time"

	"github.com/bow/courier/api"
	"github.com/mmcdole/gofeed"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Feed struct {
	DBID        DBID
	Title       string
	Description sql.NullString
	FeedURL     string
	SiteURL     sql.NullString
	Subscribed  string
	Updated     sql.NullString
	Categories  jsonArrayString
	Entries     []*Entry
}

func (f *Feed) Proto() (*api.Feed, error) {
	proto := api.Feed{
		Id:          int32(f.DBID),
		Title:       f.Title,
		FeedUrl:     f.FeedURL,
		SiteUrl:     unwrapNullString(f.SiteURL),
		Categories:  []string(f.Categories),
		Description: unwrapNullString(f.Description),
	}

	var err error

	proto.SubscriptionTime, err = toProtoTime(&f.Subscribed)
	if err != nil {
		return nil, err
	}

	proto.UpdateTime, err = toProtoTime(unwrapNullString(f.Updated))
	if err != nil {
		return nil, err
	}

	for _, entry := range f.Entries {
		ep, err := entry.Proto()
		if err != nil {
			return nil, err
		}
		proto.Entries = append(proto.Entries, ep)
	}

	return &proto, nil
}

type Entry struct {
	DBID        DBID
	Title       string
	IsRead      bool
	ExtID       string
	Updated     sql.NullString
	Published   sql.NullString
	Description sql.NullString
	Content     sql.NullString
	URL         sql.NullString
}

func (e *Entry) Proto() (*api.Entry, error) {
	proto := api.Entry{
		Id:          int32(e.DBID),
		Title:       e.Title,
		IsRead:      e.IsRead,
		ExtId:       e.ExtID,
		Description: unwrapNullString(e.Description),
		Content:     unwrapNullString(e.Content),
		Url:         unwrapNullString(e.URL),
	}

	var err error

	proto.PublicationTime, err = toProtoTime(unwrapNullString(e.Published))
	if err != nil {
		return nil, err
	}

	proto.UpdateTime, err = toProtoTime(unwrapNullString(e.Updated))
	if err != nil {
		return nil, err
	}

	return &proto, nil
}

func resolveFeedUpdateTime(feed *gofeed.Feed) *time.Time {
	// Use feed value if defined.
	var latest = feed.UpdatedParsed
	if latest != nil {
		return latest
	}
	// Otherwise try to infer from entries.
	for _, entry := range feed.Items {
		etv := resolveEntryUpdateTime(entry)
		if latest == nil {
			latest = etv
		}
		if latest != nil && etv != nil {
			if etv.After(*latest) {
				latest = etv
			}
		}
	}
	return latest
}

func resolveEntryUpdateTime(entry *gofeed.Item) *time.Time {
	// Use value if defined.
	if tv := entry.UpdatedParsed; tv != nil {
		return tv
	}
	// Otherwise use published time.
	return entry.PublishedParsed
}

func resolveEntryPublishedTime(entry *gofeed.Item) *time.Time {
	// Use value if defined.
	if tv := entry.PublishedParsed; tv != nil {
		return tv
	}
	// Otherwise use update time.
	return entry.UpdatedParsed
}

func serializeTime(tv *time.Time) *string {
	if tv == nil {
		return nil
	}
	ts := tv.UTC().Format(time.RFC3339)
	return &ts
}

func DeserializeTime(v *string) (*time.Time, error) {
	if v == nil {
		return nil, nil
	}
	pv, err := time.Parse(time.RFC3339, *v)
	if err != nil {
		return nil, err
	}
	upv := pv.UTC()
	return &upv, nil
}

func toProtoTime(v *string) (*timestamppb.Timestamp, error) {
	tv, err := DeserializeTime(v)
	if err != nil {
		return nil, err
	}
	if tv == nil {
		return nil, nil
	}
	return timestamppb.New(*tv), nil
}
