package model

import "time"

// Contests table data
type Contests struct {
	ID        int64
	Slug      string
	Title     string
	StartTime time.Time
}
