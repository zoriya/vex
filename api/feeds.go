package vex

import (
	"fmt"
	"net/http"
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
	Submitter   *User     `json:"submitter,omitempty"`
	AddedDate   time.Time `json:"addedDate"`
}

type FeedDao struct {
	Id            uuid.UUID
	Name          string
	Link          string
	FaviconUrl    string `db:"favicon_url"`
	Tags          pq.StringArray
	SubmitterId   uuid.UUID `db:"submitter_id"`
	Submitter     *User     `db:"submitter"`
	AddedDate     time.Time `db:"added_date"`
	Etag          string
	LastFetchDate *time.Time `db:"last_fetch_date"`
}

func (f *FeedDao) ToFeed() Feed {
	return Feed{
		Id:          f.Id,
		Name:        f.Name,
		Link:        f.Link,
		FaviconUrl:  f.FaviconUrl,
		Tags:        f.Tags,
		SubmitterId: f.SubmitterId,
		Submitter:   f.Submitter,
		AddedDate:   f.AddedDate,
	}
}

func (f *Feed) ToDao() FeedDao {
	return FeedDao{
		Id:          f.Id,
		Name:        f.Name,
		Link:        f.Link,
		FaviconUrl:  f.FaviconUrl,
		Tags:        f.Tags,
		SubmitterId: f.SubmitterId,
		Submitter:   f.Submitter,
		AddedDate:   f.AddedDate,
	}
}

type FeedService struct {
	database *sqlx.DB
	reader   Reader
}

func NewFeedService(db *sqlx.DB) FeedService {
	return FeedService{
		database: db,
		reader:   NewRssReader(http.DefaultClient),
	}
}

func (s FeedService) GetFeedData(link string) ([]Feed, error) {
	parsed, err := s.reader.ReadFeed(link, "", nil)
	if err != nil {
		return nil, err
	}
	return []Feed{
		{
			Id:         uuid.New(),
			Name:       parsed.Title,
			Link:       link,
			FaviconUrl: fmt.Sprintf("%s/favicon.ico", parsed.Link),
			AddedDate:  time.Now(),
		},
	}, nil
}

func (s FeedService) AddFeed(feed Feed) (Feed, error) {
	_, err := s.database.NamedExec(
		`insert into feeds (id, name, link, favicon_url, tags, submitter_id, added_date, etag, last_fetch_date)
		values (:id, :name, :link, :favicon_url, :tags, :submitter_id, :added_date, :etag, :last_fetch_date)`,
		feed.ToDao(),
	)
	if err != nil {
		return Feed{}, err
	}
	return feed, nil
}

func (s FeedService) ListFeeds() ([]Feed, error) {
	ret := []FeedDao{}
	err := s.database.Select(
		&ret,
		`select f.*, u.id as "submitter.id", u.name as "submitter.name", u.email as "submitter.email", u.password as "submitter.password"
		from feeds as f left
		join users as u on u.id = f.submitter_id
		order by added_date`,
	)
	if err != nil {
		return nil, err
	}
	return Map(ret, func(f FeedDao, _ int) Feed { return f.ToFeed() }), nil
}
