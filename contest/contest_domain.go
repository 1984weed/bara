package contest

import (
	"bara/model"
	"time"
)

// ContestWithProblem ...
type ContestWithProblem struct {
	ID        int64
	Title     string
	Slug      string
	StartTime time.Time
	Problems  []model.Problems
}

// NewContest ...
type NewContest struct {
	Title      string
	Slug       string
	StartTime  time.Time
	ProblemIDs []int64
}

// ContestRanking ...
type ContestRanking struct {
	ID        int64
	UserID    int64
	ContestID int64
	Ranking   int
}
