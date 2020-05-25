package contest

import (
	"bara/contest/domain"
	"bara/model"
	problem_domain "bara/problem/domain"
)

// Repository represent the problem's store
type Repository interface {
	GetContests(limit, offset int) ([]model.Contests, error)
	GetContest(slug string) (*model.Contests, error)
	GetContestProblems(slug string) ([]model.Problems, error)
	GetContestProblemResult(contestSlug string, problemSlug string) ([]model.ContestUserProblemSuccess, error)
	GetContestProblemsUserResults(contestID model.ContestID, userID int64) ([]model.ContestUserProblemSuccess, error)
	UpdateContestRanking(ranking []domain.ContestRanking) error
	CreateContest(newContest *domain.NewContest) (*model.Contests, error)
	RegisterContestProblem(contestProblems []domain.ContestProblemID) error
	DeleteContestProblem(contestID model.ContestID) error
	UpdateContest(contestID model.ContestID, contest *domain.NewContest) (*model.Contests, error)
	DeleteContest(slug string) error
	CreateSubmitResult(result *problem_domain.CodeResult, contestSlug string, problemSlug string, userID int64) error
}
