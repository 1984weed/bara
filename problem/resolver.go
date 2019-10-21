package problem

import (
	"bara/graphql_model"
	"context"
)

// Resolver represent the problem's resolver interface
type Resolver interface {
	GetBySlug(ctx context.Context, slug string) (*graphql_model.Question, error)
}
