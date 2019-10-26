package user

import "context"

type Resolver interface {
	Register(ctx context.Context, userID string, email string, password string) error
}
