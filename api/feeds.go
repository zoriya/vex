package vex

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Feed struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Link        string    `json:"link"`
	FaviconUrl  string    `json:"faviconUrl"`
	Tags        []string  `json:"tags"`
	SubmitterId uuid.UUID `json:"submitterId"`
	Submitter   User      `json:"submitter,omitempty"`
	AddedDate   time.Time `json:"addedDate"`
}

type FeedDao struct {
	Id          uuid.UUID
	Name        string
	Link        string
	FaviconUrl  string `db:"favicon_url"`
	Tags        pq.StringArray
	SubmitterId uuid.UUID `db:"submitter_id"`
	Submitter   User      `db:"submitter"`
	AddedDate   time.Time `db:"added_date"`
}

func (f *FeedDao) ToFeed() Feed {
	return Feed{
		Id:          f.Id,
		Name:        f.Name,
		Link:        f.Name,
		FaviconUrl:  f.FaviconUrl,
		Tags:        f.Tags,
		SubmitterId: f.SubmitterId,
		Submitter:   f.Submitter,
		AddedDate:   f.AddedDate,
	}
}

type FeedService struct {
	database *sqlx.DB
}

func NewFeedService(db *sqlx.DB) FeedService {
	return FeedService{database: db}
}

func (s FeedService) AddFeed(link string, tags []string, submitter uuid.UUID) (Feed, error) {
	feed := FeedDao{
		Id:          uuid.New(),
		Name:        link,
		Link:        link,
		FaviconUrl:  link,
		Tags:        tags,
		SubmitterId: submitter,
	}

	_, err := s.database.NamedExec(
		`insert into feeds (id, name, link, favicon_url, tags, submitter_id, added_date)
		values (:id, :name, :link, :favicon_url, :tags, :submitter_id, :added_date)`,
		feed,
	)
	if err != nil {
		return Feed{}, err
	}
	return feed.ToFeed(), nil
}

func (s FeedService) ListFeeds() ([]Feed, error) {
	ret := []FeedDao{}
	err := s.database.Select(
		&ret,
		`select f.*, u.id as "submitter.id", u.name as "submitter.name", u.email as "submitter.email", u.password as "submitter.password" from feeds
		as f left join users as u on u.id = f.submitter_id
		order by added_date`,
	)
	if err != nil {
		return nil, err
	}
	return Map(ret, func(f FeedDao, _ int) Feed { return f.ToFeed() }), nil
}