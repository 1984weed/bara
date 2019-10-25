package usecase

import (
	"bara/model"
	"bara/problem"
	"bara/problem/domain"
	"context"
	"time"

	"github.com/go-pg/pg/v9"
)

type problemUsecase struct {
	problemRepo      problem.Repository
	runInTransaction func(fn func(*pg.Tx) error) error
	contextTimeout   time.Duration
}

// NewProblemUsecase creates new a problemUsecase object of problem.Usecase interface
func NewProblemUsecase(p problem.Repository, timeout time.Duration, runInTransaction func(fn func(*pg.Tx) error) error) problem.Usecase {
	return &problemUsecase{problemRepo: p, runInTransaction: runInTransaction, contextTimeout: timeout}
}

func (p *problemUsecase) GetBySlug(ctx context.Context, slug string) (*domain.Problem, error) {
	problem, err := p.problemRepo.GetBySlug(ctx, slug)

	if err != nil {
		return nil, err
	}

	args := make([]domain.ProblemArgs, len(problem.Args))
	for i, arg := range problem.Args {
		args[i] = domain.ProblemArgs{
			Name:    arg.Name,
			VarType: arg.VarType,
		}
	}

	return &domain.Problem{
		Slug:          problem.Slug,
		Title:         problem.Title,
		Description:   problem.Description,
		LanguageSlugs: []model.CodeLanguageSlug{model.JavaScript},
		FunctionName:  problem.FunctionName,
		ProblemArgs:   args,
		OutputType:    problem.OutputType,
	}, nil
}

func (p *problemUsecase) CreateProblem(ctx context.Context, inputProblem *domain.NewProblem) (*domain.Problem, error) {
	newProblem := &model.Problems{
		Title:        inputProblem.Title,
		Slug:         inputProblem.Slug,
		Description:  inputProblem.Description,
		FunctionName: inputProblem.FunctionName,
		LanguageID:   1,
		OutputType:   inputProblem.OutputType,
		AuthorID:     0,
	}

	err := p.runInTransaction(func(tx *pg.Tx) error {
		return p.problemRepo.SaveProblem(ctx, tx, newProblem)
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
}
