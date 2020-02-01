package contest

import (
	contest "bara/contest/usecase"
	"bara/graphql_model"
	"context"
)

// Resolver represent the contest's resolver interface
type Resolver interface {
	GetContests(ctx context.Context, limit int, offset int) ([]*graphql_model.Contest, error)
}

type contestResolver struct {
	uc contest.Usecase
}

// NewContestResolver initializes the contest/ resources graphql resolver
func NewContestResolver(uc contest.Usecase) Resolver {
	return &contestResolver{uc}
}

func (cr *contestResolver) GetContests(ctx context.Context, limit int, offset int) ([]*graphql_model.Contest, error) {
	cr.GetContests(ctx, limit, offset)
	return []*graphql_model.Contest{}, nil
}
