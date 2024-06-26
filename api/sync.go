package vex

import (
	"cmp"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
)

type SyncService struct {
	reader  *Reader
	feeds   *FeedService
	entries *EntryService
}

func NewSyncService(reader *Reader, feeds *FeedService, entries *EntryService) SyncService {
	return SyncService{
		reader:  reader,
		feeds:   feeds,
		entries: entries,
	}
}

func (s SyncService) SyncFeed(feed Feed) error {
	info, err := s.reader.ReadFeed(feed.Link, feed.etag, feed.lastFetchDate)
	if err != nil {
		return err
	}
	if info == nil {
		log.Printf("Feed %v is uptodate", feed.Link)
		return nil
	}
	log.Printf("Adding %v new entries from %v", len(info.Items), feed.Link)
	entries := Map(info.Items, func(item *gofeed.Item, _ int) EntryDao {
		var date time.Time
		if item.PublishedParsed != nil {
			date = *item.PublishedParsed
		} else {
			date = time.Now()
		}

		return EntryDao{
			Id:      uuid.New(),
			Title:   item.Title,
			Link:    item.Link,
			Date:    date,
			Authors: Map(item.Authors, func(author *gofeed.Person, _ int) string { return author.Name }),
			Content: cmp.Or(item.Content, item.Description),
			FeedId:  feed.Id,
		}
	})
	err = s.entries.Add(entries)
	if err != nil {
		return err
	}
	return s.feeds.UpdateSyncStatus(feed.Id, info.ETag, &info.LastModified)
}

func (s SyncService) SyncFeeds() error {
	feeds, err := s.feeds.ListFeeds()
	if err != nil {
		log.Printf("Could not retrive feeds: %v", err)
		return err
	}
	for _, feed := range feeds {
		err := s.SyncFeed(feed)
		if err != nil {
			log.Printf("Could not sync feed %v: %v", feed.Link, err)
			s.feeds.SaveSyncError(feed.Id, err)
		}
	}
	return nil
}

func (s SyncService) SyncFeedsForever() {
	for {
		s.SyncFeeds()
		time.Sleep(15 * time.Minute)
	}
}
