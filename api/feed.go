package vex

import "github.com/google/uuid"

type Feed struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Link        string    `json:"link"`
	FaviconUrl  string    `json:"faviconUrl"`
	Tags        []string  `json:"tags"`
	SubmitterId uuid.UUID `json:"submitterId"`
}
