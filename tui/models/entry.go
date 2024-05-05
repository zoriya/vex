package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

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
