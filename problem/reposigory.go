package problem

import (
	"bara/model"
	"context"
)

// Repository represent the article's repository contract
type Repository interface {
	GetBySlug(ctx context.Context, slug string) (*model.ProblemsWithArgs, error)
	GetProblemArgsByID(ctx context.Context, problemID int64) ([]model.ProblemArgs, error)
}
