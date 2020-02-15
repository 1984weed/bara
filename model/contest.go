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

// ContestProblemUserResults table data
type ContestProblemUserResults struct {
	ID        int64
	ContestID ContestID
	ProblemID int64
	UserID    int64
	Status    string
	ExecTime  int
}

// ContestUserResults table data
type ContestUserResults struct {
	ID        int64
	UserID    int64
	ContestID int64
	ranking   int
}
