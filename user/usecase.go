package user

import "context"

type Usecase interface {
	Register(ctx context.Context, userID string, email string, password string) error
}
