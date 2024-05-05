package vex

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed"
)

type Reader struct {
	feedReader *gofeed.Parser
	client     *http.Client
}

func NewRssReader(client *http.Client) Reader {
	return Reader{
		feedReader: gofeed.NewParser(),
		client:     client,
	}
}

type GoFeed struct {
	*gofeed.Feed

	ETag         string
	LastModified time.Time
}

var gmt, _ = time.LoadLocation("GMT")

func (r *Reader) ReadFeed(url string, etag string, lastModified *time.Time) (*GoFeed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Gofeed/1.0")

	if etag != "" {
		req.Header.Set("If-None-Match", etag)
	}
	if lastModified != nil {
		req.Header.Set("If-Modified-Since", lastModified.In(gmt).Format(time.RFC1123))
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp != nil {
		defer func() {
			ce := resp.Body.Close()
			if ce != nil {
				err = ce
			}
		}()
	}

	if resp.StatusCode == http.StatusNotModified {
		return nil, nil
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, gofeed.HTTPError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
		}
	}

	feed := &GoFeed{}

	feedBody, err := r.feedReader.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	feed.Feed = feedBody

	if eTag := resp.Header.Get("Etag"); eTag != "" {
		feed.ETag = eTag
	}

	if lastModified := resp.Header.Get("Last-Modified"); lastModified != "" {
		parsed, err := time.ParseInLocation(time.RFC1123, lastModified, gmt)
		if err == nil {
			feed.LastModified = parsed
		}
	}

	return feed, nil
}
