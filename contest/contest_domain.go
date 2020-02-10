package contest

import (
	"bara/model"
	"time"
)

type ContestWithProblem struct {
	ID          int64
	Title       string
	Slug        string
	StartTime   time.Time
	ContestSlug string
	Problems    []model.Problems
}
