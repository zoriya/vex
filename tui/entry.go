package main

import (
	"time"
)

type Feed struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Url        string   `json:"url"`
	FaviconUrl string   `json:"faviconUrl"`
	Tags       []string `json:"tags"`
}

type Entry struct {
	Id           string    `json:"id"`
	ArticleTitle string    `json:"title"`
	Content      string    `json:"content"`
	Link         string    `json:"link"`
	Date         time.Time `json:"time"`

	Author       *string `json:"author"` // author not always specified
	IsRead       bool    `json:"isRead"`
	IsBookmarked bool    `json:"IsBookmarked"`
	IsIgnored    bool    `json:"isIgnored"`
	IsReadLater  bool    `json:"isReadLater"`
	Feed         Feed    `json:"feed"`
}
