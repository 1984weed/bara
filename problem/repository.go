package problem

import (
	"bara/model"
	"context"
)

// Repository represent the problem's store
type Repository interface {
	GetBySlug(ctx context.Context, slug string) (*model.ProblemsWithArgs, error)
	GetTestcaseByProblemID(ctx context.Context, problemID int64) ([]model.ProblemTestcases, error)
	SaveProblem(ctx context.Context, problem *model.Problems) error
	SaveProblemArgs(ctx context.Context, args *model.ProblemArgs) error
	SaveProblemTestcase(ctx context.Context, testcase *model.ProblemTestcases) error
}
