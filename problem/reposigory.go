package problem

import (
	"bara/model"
	"context"
)

// Repository represent the problem's store
type Repository interface {
	GetBySlug(ctx context.Context, slug string) (*model.ProblemsWithArgs, error)
}
