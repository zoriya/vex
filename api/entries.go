package vex

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Entry struct {
	Id           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Link         string    `json:"link"`
	Date         time.Time `json:"date"`
	Content      string    `json:"content"`
	Authors      []string  `json:"authors"`
	FeedId       uuid.UUID `json:"feedId"`
	Feed         Feed      `json:"feed,omitempty"`
	IsRead       bool      `json:"isRead"`
	IsBookmarked bool      `json:"isBookmarked"`
	IsReadLater  bool      `json:"isReadLater"`
	IsIgnored    bool      `json:"isIgnored"`
}

type EntryDao struct {
	Id      uuid.UUID
	Title   string
	Link    string
	Date    time.Time
	Content string
	Authors pq.StringArray
	FeedId  uuid.UUID `db:"feed_id"`
	Feed    FeedDao   `db:"feed"`

	EntryId      *uuid.UUID `db:"entry_id"`
	UserId       *uuid.UUID `db:"user_id"`
	IsRead       *bool      `db:"is_read"`
	IsBookmarked *bool      `db:"is_bookmarked"`
	IsReadLater  *bool      `db:"is_read_later"`
	IsIgnored    *bool      `db:"is_ignored"`
}

func OrDefault[T any](val *T) T {
	if val != nil {
		return *val
	}
	var ret T
	return ret
}

func (e *EntryDao) ToEntry() Entry {
	return Entry{
		Id:           e.Id,
		Title:        e.Title,
		Link:         e.Link,
		Date:         e.Date,
		Content:      e.Content,
		Authors:      e.Authors,
		FeedId:       e.FeedId,
		Feed:         e.Feed.ToFeed(),
		IsRead:       OrDefault(e.IsRead),
		IsBookmarked: OrDefault(e.IsBookmarked),
		IsReadLater:  OrDefault(e.IsReadLater),
		IsIgnored:    OrDefault(e.IsIgnored),
	}
}

type EntryService struct {
	database *sqlx.DB
}

func NewEntryService(db *sqlx.DB) EntryService {
	return EntryService{database: db}
}

func (s EntryService) Add(entries []EntryDao) error {
	_, err := s.database.NamedExec(
		`insert into entries (id, title, link, date, content, authors, feed_id)
		values (:id, :title, :link, :date, :content, :authors, :feed_id)`,
		entries,
	)
	return err
}

func (s EntryService) GetEntry(id uuid.UUID, userId uuid.UUID) (Entry, error) {
	var ret EntryDao
	err := s.database.Get(
		&ret,
		`select e.*, s.*,
		f.id as "feed.id", f.name as "feed.name", f.link as "feed.link", f.favicon_url as "feed.favicon_url",
		f.tags as "feed.tags", f.submitter_id as "feed.submitter_id", f.added_date as "feed.added_date"
		from entries as e
		left join entries_users as s on s.entry_id = e.id and s.user_id = $1
		left join feeds as f on f.id = e.feed_id
		where e.id = $2`,
		userId,
		id,
	)
	if err != nil {
		return Entry{}, err
	}
	return ret.ToEntry(), nil
}

func (s EntryService) ListEntries(userId uuid.UUID) ([]Entry, error) {
	ret := []EntryDao{}
	err := s.database.Select(
		&ret,
		`select e.*, s.*,
		f.id as "feed.id", f.name as "feed.name", f.link as "feed.link", f.favicon_url as "feed.favicon_url",
		f.tags as "feed.tags", f.submitter_id as "feed.submitter_id", f.added_date as "feed.added_date"
		from entries as e
		left join entries_users as s on s.entry_id = e.id and s.user_id = $1
		left join feeds as f on f.id = e.feed_id
		order by e.date desc`,
		userId,
	)
	if err != nil {
		return nil, err
	}
	return Map(ret, func(e EntryDao, _ int) Entry { return e.ToEntry() }), nil
}

type ChangeStatusDao struct {
	Id           uuid.UUID
	User         uuid.UUID
	IsRead       bool `db:"is_read"`
	IsBookmarked bool `db:"is_bookmarked"`
	IsReadLater  bool `db:"is_read_later"`
	IsIgnored    bool `db:"is_ignored"`
}

func (s EntryService) ChangeStatus(status ChangeStatusDao) error {
	_, err := s.database.NamedExec(
		`insert into entries_users (entry_id, user_id, is_read, is_bookmarked, is_read_later, is_ignored)
		values (:id, :user, :is_read, :is_bookmarked, :is_read_later, :is_ignored)
		on conflict(entry_id, user_id) do update set is_read = :is_read, is_bookmarked = :is_bookmarked, is_read_later = :is_read_later, is_ignored = :is_ignored`,
		status,
	)
	return err
}
