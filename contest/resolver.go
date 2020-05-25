package contest

import (
	"bara/graphql_model"
	"context"
)

// Resolver represent the contest's resolver interface
type Resolver interface {
	GetContests(ctx context.Context, limit int, offset int) ([]*graphql_model.Contest, error)
	GetContest(ctx context.Context, slug string) (*graphql_model.Contest, error)
	UpdateRankingContest(ctx context.Context, slug string) (*graphql_model.Ranking, error)
	CreateContest(ctx context.Context, contest graphql_model.NewContest) (*graphql_model.Contest, error)
	UpdateContest(ctx context.Context, contestID string, contest graphql_model.NewContest) (*graphql_model.Contest, error)
	DeleteContest(ctx context.Context, slug string) error
}
