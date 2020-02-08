package contest

import "bara/model"

type ContestWithProblem struct {
	ID          int64
	ContestSlug string
	Problems    []model.Problems
}
