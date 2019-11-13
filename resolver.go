package bara

import (
	"bara/generated"
	"bara/graphql_model"
	"bara/problem"
	"bara/user"
	"context"

	pg "github.com/go-pg/pg/v9"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	DB              *pg.DB
	ProblemResolver problem.Resolver
	UserResolver    user.Resolver
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Problems(ctx context.Context, limit *int, offset *int) ([]*graphql_model.Problem, error) {
	return r.ProblemResolver.GetProblems(ctx, *limit, *offset)
}

func (r *queryResolver) Problem(ctx context.Context, slug *string) (*graphql_model.Problem, error) {
	return r.ProblemResolver.GetBySlug(ctx, *slug)
}

func (r *queryResolver) TestNewProblem(ctx context.Context, input graphql_model.NewProblem) (*graphql_model.Problem, error) {
	return r.ProblemResolver.GetTestNewProblem(ctx, input)
}

func (r *queryResolver) Me(ctx context.Context) (*graphql_model.User, error) {
	return r.UserResolver.GetMe(ctx)
}

func (r *queryResolver) SubmissionList(ctx context.Context, problemSlug string, limit *int, offset *int) ([]*graphql_model.Submission, error) {
	return nil, nil
}

// Mutation ...
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) SubmitCode(ctx context.Context, input graphql_model.SubmitCode) (*graphql_model.CodeResult, error) {
	return r.ProblemResolver.SubmitProblem(ctx, input)
}

func (r *mutationResolver) CreateProblem(ctx context.Context, input graphql_model.NewProblem) (*graphql_model.Problem, error) {
	return r.ProblemResolver.CreateProblem(ctx, input)
}

func (r *mutationResolver) UpdateProblem(ctx context.Context, problemID int, input graphql_model.NewProblem) (*graphql_model.Problem, error) {
	return r.ProblemResolver.UpdateProblem(ctx, int64(problemID), input)
}

func (r *mutationResolver) RegisterUser(ctx context.Context, email *string, userName *string, password *string) (*graphql_model.User, error) {
	return nil, r.UserResolver.Register(ctx, *email, *userName, *password)
}
