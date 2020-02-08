package contest

import (
	"bara/model"
)

// Usecase represent the problem's usecases
type Usecase interface {
	GetContests(limit, offset int) ([]model.Contests, error)
	GetContest(contestSlug string) (*ContestWithProblem, error)
}

type contestUsecase struct {
	runner RepositoryRunner
}

// NewContestUsecase creates new a contestUsecase object of contest.Usecase interface
func NewContestUsecase(runner RepositoryRunner) Usecase {
	return &contestUsecase{runner}
}

// GetContests ...
func (c *contestUsecase) GetContests(limit, offset int) ([]model.Contests, error) {
	contests, err := c.runner.GetRepository().GetContests(limit, offset)

	if err != nil {
		return []model.Contests{}, err
	}

	return contests, err
}

// GetContest ...
func (c *contestUsecase) GetContest(contestSlug string) (*ContestWithProblem, error) {
	contest, err := c.runner.GetRepository().GetContest(contestSlug)

	if err != nil {
		return nil, err
	}

	problems, err := c.runner.GetRepository().GetContestProblems(contestSlug)

	return &ContestWithProblem{
		ID:          contest.ID,
		ContestSlug: contest.Slug,
		Problems:    problems,
	}, err
}
