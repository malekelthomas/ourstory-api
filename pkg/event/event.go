package event

type EventType int64

const (
	Article EventType = iota
)

type Event struct {
	Headline    string        `json:"headline"`
	Description string        `json:"description"`
	Meta        EventMetadata `json:"meta"`
	Location    Location      `json:"location"`
}

type EventMetadata struct {
	Type EventType
}
