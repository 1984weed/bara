package contest

import (
	"bara/contest/domain"
	"bara/model"
	problem_domain "bara/problem/domain"
)

// Usecase represent the problem's usecases
type Usecase interface {
	GetContests(limit, offset int) ([]model.Contests, error)
	GetContest(contestSlug string) (*domain.ContestWithProblem, error)
	CreateContest(newcontest *domain.NewContest) (*domain.ContestWithProblem, error)
	UpdateContest(id model.ContestID, contest *domain.NewContest) (*domain.ContestWithProblem, error)
	UpdateRankingContest(contestSlug string) error
	DeleteContest(slug string) error
	RegisterProblemResult(result *problem_domain.CodeResult, contestSlug string, problemSlug string, userID int64) error
	GetContestProblemResult(contestID model.ContestID, userID int64) (map[int64]domain.UserContestProblemResult, error)
}
