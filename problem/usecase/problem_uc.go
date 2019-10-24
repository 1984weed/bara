package usecase

import (
	"bara/model"
	"bara/problem"
	"bara/problem/domain"
	"context"
	"time"
)

type problemUsecase struct {
	problemRepo    problem.Repository
	contextTimeout time.Duration
}

// NewProblemUsecase creates new a problemUsecase object of problem.Usecase interface
func NewProblemUsecase(p problem.Repository, timeout time.Duration) problem.Usecase {
	return &problemUsecase{
		problemRepo:    p,
		contextTimeout: timeout,
	}
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
