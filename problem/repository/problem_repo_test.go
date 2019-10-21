package repository_test

import (
	"bara/model"
	"bara/problem/repository"
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetBySlug(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mockProblems := []model.Problems{
		{
			ID:           1,
			Slug:         "test-slug",
			Title:        "title",
			Description:  "description",
			FunctionName: "helloWorld",
			LanguageID:   0,
			OutputType:   "int",
			AuthorID:     0,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}
	rows := sqlmock.NewRows([]string{"id", "slug", "title", "description", "function_name", "language_id", "output_type", "author_id", "created_at", "update_at"}).
		AddRow(1, "test-slug", "test-title",
			"test-description", "helloWorld", 0, "int", 0, time.Now(), time.Now())

	query := "SELECT id,title,content, author_id, updated_at, created_at FROM article WHERE created_at > \\? ORDER BY created_at LIMIT \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := repository.NewProblemRepository(db)
	cursor := articleRepo.EncodeCursor(mockArticles[1].CreatedAt)
	num := int64(2)
	list, nextCursor, err := a.Fetch(context.TODO(), cursor, num)
	assert.NotEmpty(t, nextCursor)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}
