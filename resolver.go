package bara

import (
	"bara/generated"
	"bara/graphql_model"
	"bara/problem"
	"context"

	pg "github.com/go-pg/pg/v9"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	DB              *pg.DB
	ProblemResolver problem.Resolver
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Questions(ctx context.Context, limit *int, offset *int) ([]*graphql_model.Question, error) {
	return []*graphql_model.Question{}, nil
}

func (r *queryResolver) Question(ctx context.Context, slug *string) (*graphql_model.Question, error) {
	return r.ProblemResolver.GetBySlug(ctx, *slug)
}

func (r *queryResolver) TestNewQuestion(ctx context.Context, input graphql_model.NewQuestion) (*graphql_model.Question, error) {
	return r.ProblemResolver.GetTestNewProblem(ctx, input)
}

// Mutation ...
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) SubmitCode(ctx context.Context, input graphql_model.SubmitCode) (*graphql_model.CodeResult, error) {
	return r.ProblemResolver.SubmitProblem(ctx, input)
}

func (r *mutationResolver) CreateQuestion(ctx context.Context, input graphql_model.NewQuestion) (*graphql_model.Question, error) {
	return r.ProblemResolver.CreateProblem(ctx, input)
}
