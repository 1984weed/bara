package repository

import (
	"bara/model"
	"bara/problem"
	"context"

	pg "github.com/go-pg/pg/v9"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type problemRepository struct {
	Conn *pg.DB
}

// NewProblemRepository will create an object that represent the problem.Repository interface
func NewProblemRepository(Conn *pg.DB) problem.Repository {
	return &problemRepository{Conn}
}

func (r *problemRepository) GetBySlug(ctx context.Context, slug string) (*model.Problem, error) {
	problem := new(model.Problem)

	err := r.Conn.Model(problem).
		Where("slug = ?", slug).
		Select()

	if err != nil {
		return nil, err
	}

	return problem, nil
}

func (r *problemRepository) GetProblemArgsByID(ctx context.Context, problemID int64) ([]model.ProblemArgs, error) {
	args := new([]model.ProblemArgs)

	err := r.Conn.Model(args).
		Where("question_args.question_id = ?", problemID).
		Select()

	if err != nil {
		return nil, err
	}

	return *args, nil
}
