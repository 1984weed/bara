package mocks

import (
	"bara/contest/domain"
	"bara/model"

	problem_domain "bara/problem/domain"

	"github.com/stretchr/testify/mock"
)

// ContestUsecase mock
type ContestUsecase struct {
	mock.Mock
}

// GetBySlug mock
func (c *ContestUsecase) GetContests(limit, offset int) ([]model.Contests, error) {
	return nil, nil
}
func (c *ContestUsecase) GetContest(contestSlug string) (*domain.ContestWithProblem, error) {
	return nil, nil
}

func (c *ContestUsecase) CreateContest(newcontest *domain.NewContest) (*domain.ContestWithProblem, error) {
	return nil, nil
}
func (c *ContestUsecase) UpdateContest(id model.ContestID, contest *domain.NewContest) (*domain.ContestWithProblem, error) {
	return nil, nil
}

func (c *ContestUsecase) UpdateRankingContest(contestSlug string) error {
	return nil
}
func (c *ContestUsecase) DeleteContest(slug string) error {
	return nil
}
func (c *ContestUsecase) RegisterProblemResult(result *problem_domain.CodeResult, contestSlug string, problemSlug string, userID int64) error {
	return nil
}
func (c *ContestUsecase) GetContestProblemResult(contestID model.ContestID, userID int64) (map[int64]domain.UserContestProblemResult, error) {
	return nil, nil
}
