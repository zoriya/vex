package models

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/google/uuid"
)

type Feed struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Url        string    `json:"url"`
	FaviconUrl string    `json:"faviconUrl"`
	Tags       []string  `json:"tags"`
}

type Entry struct {
	Id           uuid.UUID `json:"id"`
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

func (e Entry) FilterValue() string {
	return e.ArticleTitle
}

func (e Entry) Title() string {
	return e.ArticleTitle
}

func (e Entry) Description() string {
	return fmt.Sprintf("%s", "my desc") // TODO: real description (tags and author + date ?)
}

type ListKeyMap struct {
	Query           key.Binding
	BookmarkToggle  key.Binding
	ReadToggle      key.Binding
	ReadLaterToggle key.Binding
	IgnoreToggle    key.Binding
	PreviewPost     key.Binding
}

func NewListKeyMap() *ListKeyMap {
	return &ListKeyMap{
		PreviewPost: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "preview post"),
		),
		Query: key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "query posts"),
		),
		BookmarkToggle: key.NewBinding(
			key.WithKeys("b"),
			key.WithHelp("b", "toggle bookmarked"),
		),
		ReadToggle: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "toggle mark as read"),
		),

		IgnoreToggle: key.NewBinding(
			key.WithKeys("x", "d"),
			key.WithHelp("x", "ignore post"),
			key.WithHelp("d", "ignore post"),
		),
		ReadLaterToggle: key.NewBinding(
			key.WithKeys("m"),
			key.WithHelp("m", "add to read later"),
		),
	}
}
