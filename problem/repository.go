package problem

import (
	"bara/model"
	"context"
)

// Repository represent the problem's store
type Repository interface {
	GetProblems(ctx context.Context, limit, offset int) ([]model.Problems, error)
	GetBySlug(ctx context.Context, slug string) (*model.ProblemsWithArgs, error)
	GetTestcaseByProblemID(ctx context.Context, problemID int64) ([]model.ProblemTestcases, error)
	SaveProblem(ctx context.Context, problem *model.Problems) error
	SaveProblemArgs(ctx context.Context, args *model.ProblemArgs) error
	SaveProblemResult(ctx context.Context, result *model.ProblemUserResults) error
	GetProblemResult(ctx context.Context, problemSlug string, userID int64, limit, offset int) ([]model.ProblemUserResults, error)
	DeleteProblemArgs(ctx context.Context, args *model.ProblemArgs) error
	SaveProblemTestcase(ctx context.Context, testcase *model.ProblemTestcases) error
	DeleteProblemTestcase(ctx context.Context, testcase *model.ProblemTestcases) error
}
