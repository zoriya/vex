package main

import (
	"time"
)

type Feed struct {
	id         string
	name       string
	url        string
	faviconUrl string
	tags       []string
}

type Entry struct {
	id      string
	title   string
	content string
	link    string
	date    time.Time

	author       *string // author not always specified
	isRead       bool
	isBookmarked bool
	isIgnored    bool
	isReadLater  bool
	feed         Feed
}
