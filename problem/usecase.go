package problem

import (
	"bara/problem/domain"
	"context"
)

// Usecase represent the problem's usecases
type Usecase interface {
	GetBySlug(ctx context.Context, slug string) (*domain.Problem, error)
	CreateProblem(ctx context.Context, input *domain.NewProblem) (*domain.Problem, error)
}
