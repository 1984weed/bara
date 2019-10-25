package problem

import (
	"bara/model"
	"context"

	"github.com/go-pg/pg/v9"
)

// Repository represent the problem's store
type Repository interface {
	GetBySlug(ctx context.Context, slug string) (*model.ProblemsWithArgs, error)
	// SaveProblemTran(ctx *pg.Tx, problem *model.Problems) error
	// GetProblemArgsByID(ctx context.Context, tx *pg.Tx, problemID int64) ([]model.ProblemArgs, error)
	GetProblemArgsByID(ctx context.Context, tx *pg.Tx, problemID int64) ([]model.ProblemArgs, error)
	// func (db *baseDB) RunInTransaction(fn func(*Tx) error) error {
	// CreateProblem(ctx context.Context, new problem *domain.NewProblem) (*model.ProblemsWithArgs, error)
	SaveProblem(ctx context.Context, tx *pg.Tx, problem *model.Problems) error
}
