package usecase

import (
	"bara/model"
	"bara/problem"
	"bara/problem/domain"
	"bara/problem/executor"
	"context"
	"time"
)

type problemUsecase struct {
	runner         problem.RepositoryRunner
	codeExecutor   executor.Client
	contextTimeout time.Duration
}

// NewProblemUsecase creates new a problemUsecase object of problem.Usecase interface
func NewProblemUsecase(runner problem.RepositoryRunner, codeExecutor executor.Client, contextTimeout time.Duration) problem.Usecase {
	return &problemUsecase{runner, codeExecutor, contextTimeout}
}

func (p *problemUsecase) GetProblems(ctx context.Context, limit, offset int) ([]domain.Problem, error) {
	rep := p.runner.GetRepository()

	problems, err := rep.GetProblems(ctx, limit, offset)

	if err != nil {
		return nil, err
	}

	modelProblems := make([]domain.Problem, len(problems))
	for i, p := range problems {
		modelProblems[i] = *domain.ConvertProblemFromTableModel(p)
	}

	return modelProblems, nil
}

func (p *problemUsecase) GetBySlug(ctx context.Context, slug string) (*domain.Problem, error) {
	rep := p.runner.GetRepository()
	problem, err := rep.GetBySlug(ctx, slug)

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

	problemTestcases := make([]domain.Testcase, len(problem.Testcases))

	for i, test := range problem.Testcases {
		problemTestcases[i] = domain.Testcase{
			Input:  test.InputText,
			Output: test.OutputText,
		}
	}

	return &domain.Problem{
		ProblemID:        problem.ID,
		Slug:             problem.Slug,
		Title:            problem.Title,
		Description:      problem.Description,
		LanguageSlugs:    []model.CodeLanguageSlug{model.JavaScript},
		FunctionName:     problem.FunctionName,
		ProblemArgs:      args,
		ProblemTestcases: problemTestcases,
		OutputType:       problem.OutputType,
	}, nil
}

func (p *problemUsecase) CreateProblem(ctx context.Context, inputProblem *domain.NewProblem) (*domain.Problem, error) {
	newProblem := &model.Problems{
		Title:        inputProblem.Title,
		Slug:         inputProblem.GetSlug(),
		Description:  inputProblem.Description,
		FunctionName: inputProblem.FunctionName,
		OutputType:   inputProblem.OutputType,
		AuthorID:     0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err := p.runner.RunInTransaction(func(repo problem.Repository) error {
		err := repo.SaveProblem(ctx, newProblem)

		if err != nil {
			return err
		}
		for i, arg := range inputProblem.ProblemArgs {
			err = repo.SaveProblemArgs(ctx, &model.ProblemArgs{
				ProblemID: newProblem.ID,
				OrderNo:   i + 1,
				Name:      arg.Name,
				VarType:   arg.VarType,
			})
			if err != nil {
				return err
			}
		}
		for _, testcase := range inputProblem.Testcases {
			err = repo.SaveProblemTestcase(ctx, &model.ProblemTestcases{
				ProblemID:  newProblem.ID,
				InputText:  testcase.GetInput(),
				OutputText: testcase.Output,
			})

			if err != nil {
				return err
			}

		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &domain.Problem{
		ProblemID:     newProblem.ID,
		Slug:          newProblem.Slug,
		Title:         newProblem.Title,
		Description:   newProblem.Description,
		LanguageSlugs: []model.CodeLanguageSlug{model.JavaScript},
		FunctionName:  newProblem.FunctionName,
		ProblemArgs:   inputProblem.ProblemArgs,
		OutputType:    newProblem.OutputType,
	}, nil
}

func (p *problemUsecase) UpdateProblem(ctx context.Context, problemID int64, inputProblem *domain.NewProblem) (*domain.Problem, error) {
	newProblem := &model.Problems{
		ID:           problemID,
		Title:        inputProblem.Title,
		Slug:         inputProblem.GetSlug(),
		Description:  inputProblem.Description,
		FunctionName: inputProblem.FunctionName,
		OutputType:   inputProblem.OutputType,
		AuthorID:     0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := p.runner.RunInTransaction(func(repo problem.Repository) error {
		err := repo.SaveProblem(ctx, newProblem)

		if err != nil {
			return err
		}
		err = repo.DeleteProblemArgs(ctx, &model.ProblemArgs{
			ProblemID: problemID,
		})
		if err != nil {
			return err
		}
		for i, arg := range inputProblem.ProblemArgs {
			err = repo.SaveProblemArgs(ctx, &model.ProblemArgs{
				ProblemID: newProblem.ID,
				OrderNo:   i + 1,
				Name:      arg.Name,
				VarType:   arg.VarType,
			})
			if err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
		err = repo.DeleteProblemTestcase(ctx, &model.ProblemTestcases{
			ProblemID: problemID,
		})
		for _, testcase := range inputProblem.Testcases {
			if err != nil {
				return err
			}

			err = repo.SaveProblemTestcase(ctx, &model.ProblemTestcases{
				ProblemID:  newProblem.ID,
				InputText:  testcase.GetInput(),
				OutputText: testcase.Output,
			})

			if err != nil {
				return err
			}

		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &domain.Problem{
		ProblemID:     newProblem.ID,
		Slug:          newProblem.Slug,
		Title:         newProblem.Title,
		Description:   newProblem.Description,
		LanguageSlugs: []model.CodeLanguageSlug{model.JavaScript},
		FunctionName:  newProblem.FunctionName,
		ProblemArgs:   inputProblem.ProblemArgs,
		OutputType:    newProblem.OutputType,
	}, nil

}

func (p *problemUsecase) SubmitProblem(ctx context.Context, code *domain.SubmitCode, userID int64) (*domain.CodeResult, error) {
	repo := p.runner.GetRepository()
	problemWithArgs, err := repo.GetBySlug(ctx, code.ProblemSlug)

	if err != nil {
		return nil, err
	}

	testcases, err := repo.GetTestcaseByProblemID(ctx, problemWithArgs.ID)

	if err != nil {
		return nil, err
	}
	domainTestcases := make([]domain.Testcase, len(testcases))

	for i, t := range testcases {
		domainTestcases[i] = domain.Testcase{
			Input:  t.InputText,
			Output: t.OutputText,
		}
	}
	testcaseStr := domain.CreateTestcase(domainTestcases)
	codeResult, err := p.codeExecutor.Exec(code.LanguageSlug, code.TypedCode, testcaseStr, problemWithArgs.FunctionName)

	if err != nil {
		return nil, err
	}

	// Record the result by the user
	err = p.runner.RunInTransaction(func(repo problem.Repository) error {
		status := codeResult.Status
		if status == "" {
			status = "error"
		}
		result := &model.ProblemUserResults{
			ProblemID:     problemWithArgs.ID,
			UserID:        userID,
			SubmittedCode: code.TypedCode,
			CodeLangID:    1,
			ExecTime:      codeResult.Time,
			Status:        status,
		}
		err = repo.SaveProblemResult(ctx, result)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return codeResult, nil
}

func (p *problemUsecase) RunProblem(ctx context.Context, code *domain.SubmitCode, inputStr string) (*domain.CodeResult, error) {
	repo := p.runner.GetRepository()
	problemWithArgs, err := repo.GetBySlug(ctx, code.ProblemSlug)

	testcase, err := repo.GetTestcaseByInput(ctx, problemWithArgs.ID, inputStr)

	if err != nil {
		return nil, err
	}

	testcaseStr := domain.CreateTestcase([]domain.Testcase{
		{
			Input:  testcase.InputText,
			Output: testcase.OutputText,
		},
	})

	codeResult, err := p.codeExecutor.Exec(code.LanguageSlug, code.TypedCode, testcaseStr, problemWithArgs.FunctionName)

	if err != nil {
		return nil, err
	}

	return codeResult, nil
}

func (p *problemUsecase) GetUsersSubmissionByProblemID(ctx context.Context, userID int64, problemSlug string, limit, offset int) ([]domain.CodeSubmission, error) {
	repo := p.runner.GetRepository()
	results, err := repo.GetProblemUserResult(ctx, problemSlug, userID, limit, offset)

	if err != nil {
		return []domain.CodeSubmission{}, err

	}

	codeSubmissions := make([]domain.CodeSubmission, len(results))

	for i, r := range results {
		codeSubmissions[i] = domain.CodeSubmission{
			ID:           r.ID,
			StatusSlug:   r.Status,
			CodeLangSlug: r.CodeLangSlug,
			ExecTime:     r.ExecTime,
			Timestamp:    r.CreatedAt,
		}
	}

	return codeSubmissions, nil
}
