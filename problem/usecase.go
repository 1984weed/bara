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
	SubmitProblem(ctx context.Context, code *domain.SubmitCode) (*domain.CodeResult, error)
}
