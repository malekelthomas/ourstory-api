package event

import (
	"time"

	"github.com/malekelthomas/ourstory-api/pkg/location"
)

type EventType int64

const (
	Article EventType = iota
	Image
	Video
)

func (e EventType) toName() string {
	switch e {
	case Article:
		return "article"
	case Image:
		return "image"
	case Video:
		return "video"
	default:
		return "err"
	}
}

//An event uploaded by a user
type Event struct {
	Headline    string            `json:"headline"`
	Description string            `json:"description"`
	Meta        EventMetadata     `json:"meta"`
	Location    location.Location `json:"location"`
	Date        time.Time         `json:"date"`
}

type EventMetadata struct {
	Type        EventType `json:"type"`
	Link        []string  `json:"link"`
	SubmittedBy string    `json:"submitted_by"`
}
