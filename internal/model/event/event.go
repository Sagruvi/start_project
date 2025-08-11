package model_event

import "time"

type Event struct {
	ID          int
	Header      string
	Description string
	Date        time.Time
	Content     Content
}

type Content struct {
	Pictures []string
	Text     string
}
