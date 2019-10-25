package repository

import (
	"bara/model"
	"bara/problem"
	"context"
	"database/sql"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type Store struct {
	// this is the actual connection(pool) to the db which has the Begin() method
	db       *pg.DB
	executor executor
}

type problemRepository struct {
	Conn   *pg.DB
	Runner func(tx interface{}) orm.DB
}

// NewProblemRepository will create an object that represent the problem.Repository interface
func NewProblemRepository(Conn *pg.DB, Runner func(tx interface{}) orm.DB) problem.Repository {
	return &problemRepository{Conn, Runner}
}

func (r *problemRepository) GetBySlug(ctx context.Context, slug string) (*model.ProblemsWithArgs, error) {
	var problem model.Problems

	err := r.Conn.Model(&problem).
		Where("problems.slug = ?", slug).
		Select()

	if err != nil {
		return nil, err
	}

	args := new([]model.ProblemArgs)
	err = r.Conn.Model(args).
		Where("problem_args.problem_id = ?", problem.ID).
		Select()

	if err != nil {
		return nil, err
	}

	return &model.ProblemsWithArgs{
		ID:           problem.ID,
		Slug:         problem.Slug,
		Title:        problem.Title,
		Description:  problem.Description,
		FunctionName: problem.FunctionName,
		OutputType:   problem.OutputType,
		AuthorID:     problem.AuthorID,
		CreatedAt:    problem.CreatedAt,
		UpdatedAt:    problem.UpdatedAt,
		Args:         *args,
	}, nil
}

func (r *problemRepository) GetProblemArgsByID(ctx context.Context, tx *pg.Tx, problemID int64) ([]model.ProblemArgs, error) {
	args := new([]model.ProblemArgs)

	err := r.Runner(tx).Model(args).
		Where("question_args.question_id = ?", problemID).
		Select()

	if err != nil {
		return nil, err
	}

	return *args, nil
}

func (r *problemRepository) SaveProblem(ctx context.Context, tx *pg.Tx, problem *model.Problems) error {
	return r.Runner(tx).Insert(problem)
}
