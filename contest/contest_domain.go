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
	UserID    int64
	ContestID model.ContestID
	Ranking   int
}

// ContestProblemID ...
type ContestProblemID struct {
	ContestID model.ContestID
	ProblemID int64
}
