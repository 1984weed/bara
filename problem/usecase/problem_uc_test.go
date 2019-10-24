package usecase_test

import (
	"bara/model"
	"bara/problem/domain"
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
	mockProblemRepo := new(mocks.ProblemRepository)

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

		u := usecase.NewProblemUsecase(mockProblemRepo, time.Second*2)

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

		u := usecase.NewProblemUsecase(mockProblemRepo, time.Second*2)

		problem, err := u.GetBySlug(context.TODO(), "test-slug")

		assert.Empty(t, problem)
		assert.Error(t, err)

		mockProblemRepo.AssertExpectations(t)
	})

}
