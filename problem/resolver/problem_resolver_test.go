package resolver_test

import (
	"bara/graphql_model"
	"bara/mocks"
	"bara/model"
	"bara/problem/domain"
	"bara/problem/resolver"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetBySlug(t *testing.T) {
	mockProblemUC := new(mocks.ProblemUsecase)
	mockContestUC := new(mocks.ContestUsecase)

	mockProblem := &domain.Problem{
		Slug:          "test-slug",
		Title:         "title",
		Description:   "description",
		FunctionName:  "helloWorld",
		LanguageSlugs: []model.CodeLanguageSlug{model.JavaScript},
		ProblemArgs: []domain.ProblemArgs{
			{Name: "target", VarType: "int"},
			{Name: "num", VarType: "int[]"},
		},
		ProblemTestcases: []domain.Testcase{
			{
				InputArray: []string{"1"},
				Input:      "input",
				Output:     "output",
			},
		},
		OutputType: "int",
	}

	t.Run("success", func(t *testing.T) {
		mockProblemUC.On("GetBySlug", mock.Anything, "test-slug").Return(mockProblem, nil).Once()

		u := resolver.NewProblemResolver(mockProblemUC, mockContestUC)

		_, err := u.GetBySlug(context.TODO(), "test-slug")

		assert.NoError(t, err)

		mockProblemUC.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockProblemUC.On("GetBySlug", mock.Anything, "test-slug").Return(nil, errors.New("Unexpexted Error")).Once()

		u := resolver.NewProblemResolver(mockProblemUC, mockContestUC)

		problem, err := u.GetBySlug(context.TODO(), "test-slug")

		assert.Empty(t, problem)
		assert.Error(t, err)

		mockProblemUC.AssertExpectations(t)
	})

}

func TestGetTestNewProblem(t *testing.T) {
	mockProblemUC := new(mocks.ProblemUsecase)
	mockContestUC := new(mocks.ContestUsecase)

	u := resolver.NewProblemResolver(mockProblemUC, mockContestUC)
	input := graphql_model.NewProblem{
		Title:        "total sum",
		Description:  "description",
		FunctionName: "helloworld",
		OutputType:   "int",
		ArgsNum:      2,
		Args: []*graphql_model.CodeArg{
			{
				Type: "int",
				Name: "target",
			},
			{
				Type: "int[]",
				Name: "num",
			},
		},
		TestCases: []*graphql_model.TestCase{},
	}

	t.Run("success", func(t *testing.T) {
		new, err := u.GetTestNewProblem(context.TODO(), input)

		assert.NoError(t, err)
		assert.Equal(t, &graphql_model.Problem{
			Title:       input.Title,
			Slug:        "total-sum",
			Description: input.Description,
			CodeSnippets: []*graphql_model.CodeSnippet{
				{
					Code: `/**
* @param {number} target
* @param {number[]} num
* @return {number} 
*/
function helloworld(target, num) {
	
}`,
					Lang: graphql_model.CodeLanguageJavaScript,
				},
			},
		}, new)
	})
}
