package repository_test

import (
	"bara/model"
	"bara/problem/repository"
	"bara/repository_suite"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type problemRepositoryTest struct {
	repository_suite.RepositoryTestSuite
}

func TestCategorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip repository test")
	}

	categorySuite := &problemRepositoryTest{
		repository_suite.RepositoryTestSuite{},
	}

	suite.Run(t, categorySuite)
}

func (a *problemRepositoryTest) SetupTest() {
	seedProblemData(a.T(), a.DB)
}

func (a *problemRepositoryTest) TearDownTest() {
	a.RepositoryTestSuite.ClearDatabase()
}

// TestGetBySlug...
func (a *problemRepositoryTest) TestGetBySlug() {
	repo := repository.NewProblemRepositoryRunner(a.DB).GetRepository()

	res, err := repo.GetBySlug(context.Background(), "test-slug")

	mockProblem := getMockProblems()[0]

	require.NoError(a.T(), err)
	assert.Equal(a.T(), mockProblem.ID, res.ID)
	assert.Equal(a.T(), mockProblem.Title, res.Title)
	assert.Equal(a.T(), mockProblem.Slug, res.Slug)
}

// TestGetBySlug...
func (a *problemRepositoryTest) TestGetTestcaseByProblemID() {
	repo := repository.NewProblemRepositoryRunner(a.DB).GetRepository()

	problems, err := repo.GetProblems(context.Background(), 0, 10)
	problemID := problems[0].ID

	res, err := repo.GetTestcaseByProblemID(context.Background(), problemID)

	testCase := getMockProblemTestCases(problemID)[0]

	require.NoError(a.T(), err)
	assert.Equal(a.T(), testCase.ProblemID, res[0].ProblemID)
	assert.Equal(a.T(), testCase.InputText, res[0].InputText)
	assert.Equal(a.T(), testCase.OutputText, res[0].OutputText)
}

// TestSaveProblem ...
func (a *problemRepositoryTest) TestSaveNewProblem() {
	repo := repository.NewProblemRepositoryRunner(a.DB).GetRepository()
	problem := &model.Problems{
		Slug:         "test-one",
		Title:        "Title one",
		Description:  "description description",
		FunctionName: "helloWorld",
		OutputType:   "int",
		AuthorID:     0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	problems, err := repo.GetProblems(context.Background(), 0, 10)

	assert.Equal(a.T(), len(getMockProblems()), len(problems))

	err = repo.SaveProblem(context.Background(), problem)

	problems, err = repo.GetProblems(context.Background(), 0, 10)
	assert.Equal(a.T(), len(getMockProblems())+1, len(problems))
	require.NoError(a.T(), err)
}

// TestSaveProblem ...
func (a *problemRepositoryTest) TestSaveUpdateProblem() {
	repo := repository.NewProblemRepositoryRunner(a.DB).GetRepository()
	// Get a problem by target slug
	oneP, _ := repo.GetBySlug(context.Background(), getMockProblems()[0].Slug)

	problem := &model.Problems{
		ID:           oneP.ID,
		Slug:         "test-one",
		Title:        "Title one",
		Description:  "description description",
		FunctionName: "helloHelloWorld",
		OutputType:   "int[]",
		AuthorID:     10,
		UpdatedAt:    time.Now(),
	}

	problems, err := repo.GetProblems(context.Background(), 0, 10)
	assert.Equal(a.T(), 2, len(problems))

	err = repo.SaveProblem(context.Background(), problem)
	require.NoError(a.T(), err)

	problems, err = repo.GetProblems(context.Background(), 0, 10)

	// do not increase the amount of problems
	assert.Equal(a.T(), 2, len(problems))

	// Get a problem by target slug
	updatedProblem, _ := repo.GetBySlug(context.Background(), problem.Slug)

	assert.Equal(a.T(), problem.ID, updatedProblem.ID)
	assert.Equal(a.T(), problem.Slug, updatedProblem.Slug)
	assert.Equal(a.T(), problem.Title, updatedProblem.Title)
	assert.Equal(a.T(), problem.Description, updatedProblem.Description)
	assert.Equal(a.T(), problem.FunctionName, updatedProblem.FunctionName)
	assert.Equal(a.T(), problem.OutputType, updatedProblem.OutputType)
}

// TestDeleteProblemArgs ...
func (a *problemRepositoryTest) TestDeleteProblemArgs() {
	repo := repository.NewProblemRepositoryRunner(a.DB).GetRepository()
	problems, _ := repo.GetProblems(context.Background(), 0, 10)

	args := getMockArgs(problems[0].ID)

	err := repo.DeleteProblemArgs(context.Background(), &args[0])

	require.NoError(a.T(), err)
}

func seedProblemData(t *testing.T, db *pg.DB) {
	languages := getMockLanguages()
	for _, l := range languages {
		err := db.Insert(&l)
		require.NoError(t, err)
	}
	problems := getMockProblems()

	for _, p := range problems {
		err := db.Insert(&p)
		args := getMockArgs(p.ID)
		for _, arg := range args {
			fmt.Println(arg)
			err = db.Insert(&arg)
			require.NoError(t, err)
		}
		testCases := getMockProblemTestCases(p.ID)
		for _, tc := range testCases {
			err = db.Insert(&tc)
			require.NoError(t, err)
		}

		require.NoError(t, err)
	}
}

func getMockLanguages() []model.CodeLanguages {
	return []model.CodeLanguages{
		{
			ID:   1,
			Name: "JavaScript",
			Slug: "javascript",
		},
	}
}

func getMockProblems() []model.Problems {
	return []model.Problems{
		{
			Slug:         "test-slug",
			Title:        "title",
			Description:  "description",
			FunctionName: "helloWorld",
			OutputType:   "int",
			AuthorID:     0,
		},
		{
			Slug:         "test-slug-2",
			Title:        "title-2",
			Description:  "description-2",
			FunctionName: "calcSum",
			OutputType:   "int",
			AuthorID:     0,
		},
	}
}

func getMockProblemTestCases(problemID int64) []model.ProblemTestcases {
	return []model.ProblemTestcases{{
		ProblemID:  problemID,
		InputText:  "9\n8",
		OutputText: "10",
	}}
}

func getMockArgs(problemID int64) []model.ProblemArgs {
	return []model.ProblemArgs{
		{
			ProblemID: problemID,
			OrderNo:   1,
			Name:      "target",
			VarType:   "int",
		},
		{
			ProblemID: problemID,
			OrderNo:   2,
			Name:      "nums",
			VarType:   "int[]",
		},
	}

}
