package bara

import (
	"bara/contest"
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
	ContestResolver contest.Resolver
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

func (r *queryResolver) User(ctx context.Context, userName *string) (*graphql_model.User, error) {
	return r.UserResolver.GetUser(ctx, *userName)
}

func (r *queryResolver) SubmissionList(ctx context.Context, problemSlug string, limit *int, offset *int) ([]*graphql_model.Submission, error) {
	var limitNum int
	if limit == nil {
		limitNum = 25
	} else {
		limitNum = *limit
	}
	var offsetNum int
	if offset == nil {
		offsetNum = 25
	} else {
		offsetNum = *offset
	}
	return r.ProblemResolver.GetUsersSubmissionByProblemID(ctx, problemSlug, limitNum, offsetNum)
}
func (r *queryResolver) Contests(ctx context.Context, limit *int, offset *int) ([]*graphql_model.Contest, error) {
	var limitNum int
	if limit == nil {
		limitNum = 25
	} else {
		limitNum = *limit
	}
	var offsetNum int
	if offset == nil {
		offsetNum = 25
	} else {
		offsetNum = *offset
	}
	return r.ContestResolver.GetContests(ctx, limitNum, offsetNum)
}

func (r *queryResolver) Contest(ctx context.Context, slug string) (*graphql_model.Contest, error) {
	contest, err := r.ContestResolver.GetContest(ctx, slug)

	if err != nil {
		return nil, err
	}

	return contest, nil
}

// Mutation ...
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) SubmitCode(ctx context.Context, input graphql_model.SubmitCode) (*graphql_model.CodeResult, error) {
	return r.ProblemResolver.SubmitProblem(ctx, input)
}

func (r *mutationResolver) SubmitContestCode(ctx context.Context, contestSlug string, input graphql_model.SubmitCode) (*graphql_model.CodeResult, error) {
	return r.ProblemResolver.SubmitContestCode(ctx, contestSlug, input)
}

func (r *mutationResolver) TestRunCode(ctx context.Context, inputStr string, input graphql_model.SubmitCode) (*graphql_model.CodeResult, error) {
	return r.ProblemResolver.TestRunCode(ctx, inputStr, input)
}

func (r *mutationResolver) CreateProblem(ctx context.Context, input graphql_model.NewProblem) (*graphql_model.Problem, error) {
	return r.ProblemResolver.CreateProblem(ctx, input)
}

func (r *mutationResolver) UpdateProblem(ctx context.Context, problemID int, input graphql_model.NewProblem) (*graphql_model.Problem, error) {
	return r.ProblemResolver.UpdateProblem(ctx, int64(problemID), input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input graphql_model.UserInput) (*graphql_model.User, error) {
	return r.UserResolver.UpdateUser(ctx, input)
}

func (r *mutationResolver) CreateContest(ctx context.Context, newContest graphql_model.NewContest) (*graphql_model.Contest, error) {
	return r.ContestResolver.CreateContest(ctx, newContest)
}

func (r *mutationResolver) UpdateContest(ctx context.Context, contestID string, newContest graphql_model.NewContest) (*graphql_model.Contest, error) {
	return r.ContestResolver.UpdateContest(ctx, contestID, newContest)
}

func (r *mutationResolver) DeleteContest(ctx context.Context, contestSlug string) (*bool, error) {
	err := r.ContestResolver.DeleteContest(ctx, contestSlug)
	if err != nil {
		b := false
		return &b, err
	}

	b := true
	return &b, nil
}
