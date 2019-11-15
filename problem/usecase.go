package problem

import (
	"bara/problem/domain"
	"context"
)

// Usecase represent the problem's usecases
type Usecase interface {
	GetProblems(ctx context.Context, limit, offset int) ([]domain.Problem, error)
	GetBySlug(ctx context.Context, slug string) (*domain.Problem, error)
	CreateProblem(ctx context.Context, input *domain.NewProblem) (*domain.Problem, error)
	UpdateProblem(ctx context.Context, problemID int64, input *domain.NewProblem) (*domain.Problem, error)
	SubmitProblem(ctx context.Context, code *domain.SubmitCode, userID int64) (*domain.CodeResult, error)
	GetUsersSubmissionByProblemID(ctx context.Context, userID int64, problemSlug string, limit, offset int) ([]domain.CodeSubmission, error)
}
