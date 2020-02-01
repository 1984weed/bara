package contest

import (
	"bara/problem"
	"context"
)

// Usecase represent the problem's usecases
type Usecase interface {
	GetContests(ctx context.Context, limit, offset int)
}

type contestUsecase struct {
	runner problem.RepositoryRunner
}

// NewContestUsecase creates new a contestUsecase object of contest.Usecase interface
func NewContestUsecase(runner problem.RepositoryRunner) Usecase {
	return &contestUsecase{runner}
}

func (c *contestUsecase) GetContests(ctx context.Context, limit, offset int) {
}
