package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Url         string    `json:"link"`
	FaviconUrl  string    `json:"faviconUrl"`
	Tags        []string  `json:"tags"`
	AddedDate   time.Time `json:"addedDate"`
	SubmitterId uuid.UUID `json:"submitterId"`
}

func (f Feed) FilterValue() string {
	return f.Name
}

func (f Feed) Title() string {
	return f.Name
}

func (f Feed) Description() string {
	return fmt.Sprintf("%s", "my desc") // TODO: real description (tags, submitter, error status, last sync)
}
