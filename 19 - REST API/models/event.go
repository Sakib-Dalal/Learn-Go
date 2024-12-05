package models

import "time"

type Event struct {
	ID       int // `binding:"required"`
	Name     string
	Age      string
	DateTime time.Time
	UserID   int
}

var events = []Event{}

func (e Event) Save() {
	// later: add it to database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
