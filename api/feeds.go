package vex

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Feed struct {
	Id          uuid.UUID
	Name        string
	Link        string
	FaviconUrl  string `db:"favicon_url"`
	Tags        pq.StringArray
	SubmitterId uuid.UUID `db:"submitter_id"`
}

type FeedService struct {
	database *sqlx.DB
}

func NewFeedService(db *sqlx.DB) FeedService {
	return FeedService{database: db}
}

func (s FeedService) AddFeed(link string, tags []string, submitter uuid.UUID) (Feed, error) {
	feed := Feed{
		Id:          uuid.New(),
		Name:        link,
		Link:        link,
		FaviconUrl:  link,
		Tags:        tags,
		SubmitterId: submitter,
	}

	_, err := s.database.NamedExec(
		`insert into feeds (id, name, link, favicon_url, tags, submitter_id)
		values (:id, :name, :link, :favicon_url, :tags, :submitter_id)`,
		feed,
	)
	if err != nil {
		return Feed{}, err
	}
	return feed, nil
}
