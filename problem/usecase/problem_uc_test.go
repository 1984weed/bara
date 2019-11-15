package usecase_test

import (
	"bara/model"
	"bara/problem/domain"
	"bara/problem/executor"
	"bara/problem/mocks"
	"bara/problem/usecase"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetBySlug(t *testing.T) {
	mockProblemRunner := new(mocks.RepositoryRunner)
	mockProblemRepo := new(mocks.ProblemRepository)

	mockProblemRunner.On("GetRepository").Return(mockProblemRepo)

	mockProblem := &model.ProblemsWithArgs{
		ID:           1,
		Slug:         "test-slug",
		Title:        "title",
		Description:  "description",
		FunctionName: "helloWorld",
		Args: []model.ProblemArgs{
			{
				ID:        0,
				ProblemID: 1,
				OrderNo:   1,
				Name:      "target",
				VarType:   "int",
			},
			{
				ID:        2,
				ProblemID: 1,
				OrderNo:   2,
				Name:      "num",
				VarType:   "int[]",
			},
		},
		OutputType: "int",
		AuthorID:   0,
	}

	t.Run("success", func(t *testing.T) {
		mockProblemRepo.On("GetBySlug", mock.Anything, "test-slug").Return(mockProblem, nil).Once()

		u := usecase.NewProblemUsecase(mockProblemRunner, executor.NewExecutorClient(false, time.Second*30), time.Second*2)

		problem, err := u.GetBySlug(context.TODO(), "test-slug")

		assert.NoError(t, err)
		assert.Equal(t, mockProblem.Slug, problem.Slug)
		assert.Equal(t, mockProblem.Title, problem.Title)
		assert.Equal(t, mockProblem.Description, problem.Description)
		assert.Equal(t, []model.CodeLanguageSlug{model.JavaScript}, problem.LanguageSlugs)
		assert.Equal(t, mockProblem.FunctionName, problem.FunctionName)
		assert.Equal(t, []domain.ProblemArgs{
			{Name: "target", VarType: "int"},
			{Name: "num", VarType: "int[]"},
		}, problem.ProblemArgs)
		assert.Equal(t, mockProblem.OutputType, problem.OutputType)

		mockProblemRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockProblemRepo.On("GetBySlug", mock.Anything, "test-slug").Return(nil, errors.New("Unexpexted Error")).Once()

		u := usecase.NewProblemUsecase(mockProblemRunner, executor.NewExecutorClient(false, time.Second*30), time.Second*2)

		problem, err := u.GetBySlug(context.TODO(), "test-slug")

		assert.Empty(t, problem)
		assert.Error(t, err)

		mockProblemRepo.AssertExpectations(t)
	})
}

func TestGetProblems(t *testing.T) {
	mockProblemRunner := new(mocks.RepositoryRunner)
	mockProblemRepo := new(mocks.ProblemRepository)

	mockProblemRunner.On("GetRepository").Return(mockProblemRepo)

	mockProblems := []model.Problems{
		{
			ID:           1,
			Slug:         "test-slug",
			Title:        "title",
			Description:  "description",
			FunctionName: "helloWorld",
			OutputType:   "int",
			AuthorID:     0,
		},
		{
			ID:           2,
			Slug:         "test-slug-2",
			Title:        "title-2",
			Description:  "description-2",
			FunctionName: "helloWorld-2",
			OutputType:   "int[]",
			AuthorID:     10,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockProblemRepo.On("GetProblems", mock.Anything, 0, 10).Return(mockProblems, nil).Once()

		u := usecase.NewProblemUsecase(mockProblemRunner, executor.NewExecutorClient(false, time.Second*30), time.Second*2)

		problems, err := u.GetProblems(context.TODO(), 0, 10)

		assert.NoError(t, err)
		for i, p := range problems {
			assert.Equal(t, mockProblems[i].Slug, p.Slug)
			assert.Equal(t, mockProblems[i].Title, p.Title)
			assert.Equal(t, mockProblems[i].Description, p.Description)
			assert.Equal(t, mockProblems[i].FunctionName, p.FunctionName)
			assert.Equal(t, mockProblems[i].OutputType, p.OutputType)
		}

		mockProblemRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockProblemRepo.On("GetProblems", mock.Anything, 0, 10).Return(nil, errors.New("Unexpexted Error")).Once()

		u := usecase.NewProblemUsecase(mockProblemRunner, executor.NewExecutorClient(false, time.Second*30), time.Second*2)

		problems, err := u.GetProblems(context.TODO(), 0, 10)

		assert.Equal(t, 0, len(problems))
		assert.Error(t, err)

		mockProblemRepo.AssertExpectations(t)
	})
}

func TestUpdateProblem(t *testing.T) {
	runner, repo := mocks.NewRepositoryRunnerMock()

	t.Run("success", func(t *testing.T) {
		repo.On("SaveProblem", mock.Anything, mock.Anything).Return(nil).Once()
		repo.On("DeleteProblemArgs", mock.Anything, mock.Anything).Return(nil).Twice()
		repo.On("SaveProblemArgs", mock.Anything, mock.Anything).Return(nil).Twice()
		repo.On("DeleteProblemArgs", mock.Anything, mock.Anything).Return(nil).Twice()
		repo.On("SaveProblemArgs", mock.Anything, mock.Anything).Return(nil).Twice()
		repo.On("DeleteProblemTestcase", mock.Anything, mock.Anything).Return(nil).Twice()
		repo.On("SaveProblemTestcase", mock.Anything, mock.Anything).Return(nil).Twice()

		u := usecase.NewProblemUsecase(runner, executor.NewExecutorClient(false, time.Second*30), time.Second*2)
		pArgs := []domain.ProblemArgs{
			{
				Name:    "target",
				VarType: "int",
			},
			{
				Name:    "num",
				VarType: "int[]",
			},
		}
		pTestCases := []domain.Testcase{
			{
				InputArray: []string{"5", "1,2,3,4,5"},
				Output:     "7",
			},
			{
				InputArray: []string{"99", "100,101,102,103,104"},
				Output:     "7",
			},
		}

		updatedProblem := &domain.NewProblem{
			Title:        "Updated title",
			Description:  "Updated description",
			OutputType:   "int",
			FunctionName: "helloWorld",
			ProblemArgs:  pArgs,
			Testcases:    pTestCases,
		}

		p, err := u.UpdateProblem(context.TODO(), 1, updatedProblem)

		expectProblem := &domain.Problem{
			Slug:             "updated-title",
			Title:            updatedProblem.Title,
			Description:      updatedProblem.Description,
			LanguageSlugs:    []model.CodeLanguageSlug{model.JavaScript},
			FunctionName:     updatedProblem.FunctionName,
			ProblemArgs:      updatedProblem.ProblemArgs,
			ProblemTestcases: updatedProblem.Testcases,
			OutputType:       updatedProblem.OutputType,
		}

		assert.NoError(t, err)
		assert.Equal(t, expectProblem.Title, p.Title)
		assert.Equal(t, expectProblem.Slug, p.Slug)
		assert.Equal(t, expectProblem.Description, p.Description)
		assert.Equal(t, expectProblem.FunctionName, p.FunctionName)
		assert.Equal(t, expectProblem.ProblemArgs, p.ProblemArgs)
		assert.Equal(t, []domain.Testcase(nil), p.ProblemTestcases)
	})
}

func TestGetUsersSubmissionByProblemID(t *testing.T) {
	runner, repo := mocks.NewRepositoryRunnerMock()
	t.Run("success", func(t *testing.T) {
		u := usecase.NewProblemUsecase(runner, executor.NewExecutorClient(false, time.Second*30), time.Second*2)

		problemSlug, userID, limit, offset := "problem-slug", 1, 10, 0
		mockSubmissions := []model.ProblemUserSubmission{
			{
				ID:            9999,
				SubmittedCode: "function test(){}",
				Status:        "success",
				CodeLangSlug:  model.JavaScript,
				ExecTime:      10,
				CreatedAt:     time.Now(),
			},
			{
				ID:            10000,
				SubmittedCode: "function test(){console.log('-----------')}",
				Status:        "fail",
				CodeLangSlug:  model.JavaScript,
				ExecTime:      0,
				CreatedAt:     time.Now(),
			},
		}

		repo.On("GetProblemUserResult", mock.Anything, problemSlug, int64(userID), limit, offset).Return(mockSubmissions, nil).Once()

		actual, err := u.GetUsersSubmissionByProblemID(context.TODO(), int64(userID), problemSlug, limit, offset)
		expected := make([]domain.CodeSubmission, len(mockSubmissions))

		for i, s := range mockSubmissions {
			expected[i] = domain.CodeSubmission{
				ID:           s.ID,
				StatusSlug:   s.Status,
				CodeLangSlug: s.CodeLangSlug,
				ExecTime:     s.ExecTime,
				Timestamp:    actual[i].Timestamp,
			}
		}

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}
