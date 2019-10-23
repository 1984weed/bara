package repository_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-pg/pg/v9"
	"github.com/stretchr/testify/suite"
)

type articleRepositoryTest struct {
	RepositoryTestSuite
}

func TestCategorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip category mysql repository test")
	}
	dsn := os.Getenv("MYSQL_TEST_URL")
	if dsn == "" {
		dsn = "root:root-pass@tcp(localhost:33060)/testing?parseTime=1&loc=Asia%2FJakarta&charset=utf8mb4&collation=utf8mb4_unicode_ci"
	}

	categorySuite := &articleRepositoryTest{
		RepositoryTestSuite{
			Config: &pg.Options{
				User:     "postgres",
				Password: "postgres",
				Network:  "tcp",
				Addr:     "0.0.0.0:5555",
				Database: "bara",
			},
		},
	}

	suite.Run(t, categorySuite)
}

func (s *articleRepositoryTest) TearDownTest() {
	s.RepositoryTestSuite.ClearDatabase()
}

func (m *articleRepositoryTest) TestStore() {
	fmt.Println("=========================Test====================")
}

// func TestGetBySlug(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	mockProblems := []model.Problems{
// 		{
// 			ID:           1,
// 			Slug:         "test-slug",
// 			Title:        "title",
// 			Description:  "description",
// 			FunctionName: "helloWorld",
// 			LanguageID:   0,
// 			OutputType:   "int",
// 			AuthorID:     0,
// 			CreatedAt:    time.Now(),
// 			UpdatedAt:    time.Now(),
// 		},
// 	}
// 	rows := sqlmock.NewRows([]string{"id", "slug", "title", "description", "function_name", "language_id", "output_type", "author_id", "created_at", "update_at"}).
// 		AddRow(1, "test-slug", "test-title",
// 			"test-description", "helloWorld", 0, "int", 0, time.Now(), time.Now())

// 	query := "SELECT id,title,content, author_id, updated_at, created_at FROM article WHERE created_at > \\? ORDER BY created_at LIMIT \\?"

// 	mock.ExpectQuery(query).WillReturnRows(rows)
// 	a := repository.NewProblemRepository(db)
// 	cursor := articleRepo.EncodeCursor(mockArticles[1].CreatedAt)
// 	num := int64(2)
// 	list, nextCursor, err := a.Fetch(context.TODO(), cursor, num)
// 	assert.NotEmpty(t, nextCursor)
// 	assert.NoError(t, err)
// 	assert.Len(t, list, 2)
// }
