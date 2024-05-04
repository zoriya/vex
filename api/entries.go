package vex

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Entry struct {
	Id      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Link    string    `json:"link"`
	Date    time.Time `json:"date"`
	Content string    `json:"content"`
	Author  *string   `json:"author"`
	FeedId  uuid.UUID `json:"feedId"`
	Feed    Feed      `json:"feed,omitempty"`
}

type EntryDao struct {
	Id      uuid.UUID
	Title   string
	Link    string
	Date    time.Time
	Content string
	Author  *string
	FeedId  uuid.UUID `db:"feed_id"`
	Feed    FeedDao   `db:"feed"`
}

func (e *EntryDao) ToEntry() Entry {
	return Entry{
		Id:      e.Id,
		Title:   e.Title,
		Link:    e.Link,
		Date:    e.Date,
		Content: e.Content,
		Author:  e.Author,
		FeedId:  e.FeedId,
		Feed:    e.Feed.ToFeed(),
	}
}

type EntryService struct {
	database *sqlx.DB
}

func NewEntryService(db *sqlx.DB) EntryService {
	return EntryService{database: db}
}

func (s EntryService) ListEntries() ([]Entry, error) {
	ret := []EntryDao{}
	err := s.database.Select(
		&ret,
		`select e.*, f.id as "feed.id", f.name as "feed.name", f.link as "feed.link", f.favicon_url as "feed.favicon_url",
		f.tags as "feed.tags", f.submitter_id as "feed.submitter_id", f.added_date as "feed.added_date"
		from entries as e
		left join feeds as f on f.id = e.feed_id
		order by e.date`,
	)
	if err != nil {
		return nil, err
	}
	return Map(ret, func(e EntryDao, _ int) Entry { return e.ToEntry() }), nil
}
