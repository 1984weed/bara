package problem

import (
	"bara/problem/domain"
	"context"
)

// Usecase represent the article's usecases
type Usecase interface {
	// Fetch(ctx context.Context, cursor string, num int64) ([]*models.Article, string, error)
	GetBySlug(ctx context.Context, slug string) (*domain.Problem, error)
	// GetByID(ctx context.Context, id int64) (*models.Article, error)
	// Update(ctx context.Context, ar *models.Article) error
	// GetByTitle(ctx context.Context, title string) (*models.Article, error)
	// Store(context.Context, *models.Article) error
	// Delete(ctx context.Context, id int64) error
}
