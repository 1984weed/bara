package model

import "time"

type ContestID int64

// Contests table data
type Contests struct {
	ID        ContestID
	Slug      string
	Title     string
	StartTime time.Time
}
