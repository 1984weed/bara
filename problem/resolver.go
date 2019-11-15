package problem

import (
	"bara/graphql_model"
	"context"
)

// Resolver represent the problem's resolver interface
type Resolver interface {
	GetProblems(ctx context.Context, limit int, offset int) ([]*graphql_model.Problem, error)
	GetBySlug(ctx context.Context, slug string) (*graphql_model.Problem, error)
	GetTestNewProblem(ctx context.Context, input graphql_model.NewProblem) (*graphql_model.Problem, error)
	CreateProblem(ctx context.Context, input graphql_model.NewProblem) (*graphql_model.Problem, error)
	UpdateProblem(ctx context.Context, problemID int64, input graphql_model.NewProblem) (*graphql_model.Problem, error)
	SubmitProblem(ctx context.Context, input graphql_model.SubmitCode) (*graphql_model.CodeResult, error)
	GetUsersSubmissionByProblemID(ctx context.Context, problemSlug string, limit, offset int) ([]*graphql_model.Submission, error)
}
