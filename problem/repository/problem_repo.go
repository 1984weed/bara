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
	var problems []model.Problems

	_, err := r.Conn.Query(
		&problems, `
			SELECT 
				id,
				slug,
				title,
				description,
				function_name,
				output_type,
				author_id,
				created_at,
				updated_at
			FROM problems p
			ORDER BY p.id
			LIMIT ? OFFSET ?
		`, limit, offset)

	if err != nil {
		return nil, err
	}

	return problems, err
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

func (r *problemRepository) GetTestcaseByInput(ctx context.Context, problemID int64, input string) (*model.ProblemTestcases, error) {
	qts := new(model.ProblemTestcases)

	err := r.Conn.Model(qts).
		Where("problem_testcases.problem_id = ?", problemID).
		Where("problem_testcases.input_text = ?", input).
		Select()

	if err != nil {
		return nil, err
	}

	return qts, err
}

func (r *problemRepository) SaveProblem(ctx context.Context, problem *model.Problems) error {
	_, err := r.Conn.Model(problem).
		OnConflict("(id) DO UPDATE").
		Set("title = EXCLUDED.title").
		Set("slug = EXCLUDED.slug").
		Set("description = EXCLUDED.description").
		Set("function_name = EXCLUDED.function_name").
		Set("output_type = EXCLUDED.output_type").
		Set("author_id = EXCLUDED.author_id").
		Set("updated_at = EXCLUDED.updated_at").
		Insert()

	return err
}

func (r *problemRepository) SaveProblemResult(ctx context.Context, result *model.ProblemUserResults) error {
	return r.Conn.Insert(result)
}

func (r *problemRepository) GetProblemUserResult(ctx context.Context, problemSlug string, userID int64, limit, offset int) ([]model.ProblemUserSubmission, error) {
	var problemUserResuts []model.ProblemUserSubmission

	_, err := r.Conn.Query(
		&problemUserResuts, `
			SELECT 
				pur.id as id,
				pur.submitted_code,
				pur.status,
				l.slug as code_lang_slug,
				pur.exec_time,
				pur.created_at
			FROM problem_user_results pur, problems p, code_languages l
			WHERE user_id = ?
			AND pur.problem_id = p.id 
			AND p.slug = ?
			AND l.ID = pur.code_lang_id
			ORDER BY id 
			LIMIT ? OFFSET ?
		`, userID, problemSlug, limit, offset)

	if err != nil {
		return []model.ProblemUserSubmission{}, err
	}

	return problemUserResuts, err
}

func (r *problemRepository) SaveProblemArgs(ctx context.Context, args *model.ProblemArgs) error {
	_, err := r.Conn.Model(args).Insert()

	return err
}

func (r *problemRepository) DeleteProblemArgs(ctx context.Context, args *model.ProblemArgs) error {
	_, err := r.Conn.Model(args).Where("problem_id = ?problem_id").Delete()

	return err
}

func (r *problemRepository) SaveProblemTestcase(ctx context.Context, testcase *model.ProblemTestcases) error {
	return r.Conn.Insert(testcase)
}

func (r *problemRepository) DeleteProblemTestcase(ctx context.Context, testcase *model.ProblemTestcases) error {
	_, err := r.Conn.Model(testcase).Where("problem_id = ?problem_id").Delete()
	return err
}
