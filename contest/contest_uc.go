package contest

import (
	"bara/model"
)

// Usecase represent the problem's usecases
type Usecase interface {
	GetContests(limit, offset int) ([]model.Contests, error)
}

type contestUsecase struct {
	runner RepositoryRunner
}

// NewContestUsecase creates new a contestUsecase object of contest.Usecase interface
func NewContestUsecase(runner RepositoryRunner) Usecase {
	return &contestUsecase{runner}
}

// GetContests
func (c *contestUsecase) GetContests(limit, offset int) ([]model.Contests, error) {
	contests, err := c.runner.GetRepository().GetContests(limit, offset)

	if err != nil {
		return []model.Contests{}, err
	}

	return contests, err
}
