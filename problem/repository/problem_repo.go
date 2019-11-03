package repository

import (
	"bara/model"
	"bara/problem"
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type problemRepositoryRunner struct {
	Conn *pg.DB
}

func (p *problemRepositoryRunner) RunInTransaction(fn func(r problem.Repository) error) error {
	return p.Conn.RunInTransaction(func(tx *pg.Tx) error {
		pr := newProblemRepository(interface{}(tx).(orm.DB))
		return fn(pr)
	})
}

func (p *problemRepositoryRunner) GetRepository() problem.Repository {
	return newProblemRepository(interface{}(p.Conn).(orm.DB))
}

// NewProblemRepositoryRunner will create an object that represent the problem.Repository Runner Interface
func NewProblemRepositoryRunner(Conn *pg.DB) problem.RepositoryRunner {
	return &problemRepositoryRunner{Conn}
}

type problemRepository struct {
	Conn orm.DB
}

// newProblemRepository will create an object that represent the problem.Repository interface
func newProblemRepository(Conn orm.DB) problem.Repository {
	return &problemRepository{Conn}
}
func (r *problemRepository) GetProblems(ctx context.Context, limit, offset int) ([]model.Problems, error) {
	problems := new([]model.Problems)

	err := r.Conn.Model(problems).
		Select()

	return *problems, err
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
	testcases := new([]model.ProblemTestcases)

	err = r.Conn.Model(testcases).
		Where("problem_testcases.problem_id = ?", problem.ID).
		Select()

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
		Testcases:    *testcases,
	}, nil
}

func (r *problemRepository) GetTestcaseByProblemID(ctx context.Context, problemID int64) ([]model.ProblemTestcases, error) {
	qts := new([]model.ProblemTestcases)

	err := r.Conn.Model(qts).
		Where("problem_testcases.problem_id = ?", problemID).
		Select()

	if err != nil {
		return []model.ProblemTestcases{}, err
	}

	return *qts, err
}

func (r *problemRepository) SaveProblem(ctx context.Context, problem *model.Problems) error {
	return r.Conn.Insert(problem)
}

func (r *problemRepository) SaveProblemArgs(ctx context.Context, args *model.ProblemArgs) error {
	return r.Conn.Insert(args)
}

func (r *problemRepository) SaveProblemTestcase(ctx context.Context, testcase *model.ProblemTestcases) error {
	return r.Conn.Insert(testcase)
}
