package repository_test

import (
	"bara/model"
	"bara/problem/repository"
	"bara/repository_suite"
	"context"
	"testing"

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
	repo := repository.NewProblemRepository(a.DB)

	res, err := repo.GetBySlug(context.Background(), "test-slug")

	mockProblem := getMockProblems()[0]
	require.NoError(a.T(), err)
	assert.Equal(a.T(), mockProblem.ID, res.ID)
	assert.Equal(a.T(), mockProblem.Title, res.Title)
	assert.Equal(a.T(), mockProblem.Slug, res.Slug)
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
			ID:           1,
			Slug:         "test-slug",
			Title:        "title",
			Description:  "description",
			FunctionName: "helloWorld",
			LanguageID:   1,
			OutputType:   "int",
			AuthorID:     0,
		},
		{
			ID:           2,
			Slug:         "test-slug-2",
			Title:        "title-2",
			Description:  "description-2",
			FunctionName: "calcSum",
			LanguageID:   1,
			OutputType:   "int",
			AuthorID:     0,
		},
	}
}
