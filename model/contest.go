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

// ContestProblems ...
type ContestProblems struct {
	ID        int64
	ContestID ContestID
	ProblemID int64
	OrderID   int
}

// ContestUserResults table data
type ContestUserResults struct {
	ID        int64
	UserID    int64
	ContestID ContestID
	Ranking   int
}

// ContestProblemUserTime
type ContestUserProblemSuccess struct {
	UserID    int64
	ExecTime  int
	ProblemID int64
	CreatedAt time.Time
	Status    string
}
