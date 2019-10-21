package usecase

import (
	"bara/model"
	"bara/problem"
	"bara/problem/domain"
	"context"
	"time"
	// "github.com/bxcodec/go-clean-arch/article"
	// "github.com/bxcodec/go-clean-arch/author"
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
	// type ProblemArgs struct {
	// 	Name    string
	// 	VarType string
	// }
	// args, err := p.problemRepo.GetProblemArgsByID(ctx, problem.ID)

	// if err != nil {
	// 	return nil, err
	// }

	// language, err := p.problemRepo.get

	return &domain.Problem{
		Slug:          problem.Slug,
		Title:         problem.Title,
		Description:   problem.Description,
		LanguageSlugs: []model.CodeLanguageSlug{model.JavaScript},
		FunctionName:  problem.FunctionName,
		ProblemArgs:   args,
		OutputType:    problem.OutputType,
	}, nil
	// question := new(remote.Question)

	// err := r.DB.Model(question).
	// 	Where("slug = ?", *slug).
	// 	Select()

	// if err != nil {
	// 	return nil, err
	// }

	// args := new([]remote.QuestionArgs)

	// err = r.DB.Model(args).
	// 	Where("question_args.question_id = ?", question.ID).
	// 	Select()
	// if err != nil {
	// 	return nil, err
	// }

	// return &Question{
	// 	Slug:        question.Slug,
	// 	Title:       question.Title,
	// 	Description: question.Description,
	// 	CodeSnippets: []*CodeSnippet{
	// 		{
	// 			Code: makeSnippets(question.FunctionName, args, question.OutputType),
	// 			Lang: CodeLanguageJavaScript,
	// 		},
	// 	},
	// }, nil
}
