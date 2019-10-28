package user

import "context"

type Resolver interface {
	Register(ctx context.Context, userName string, email string, password string) error
	// Login(ctx context.Context, email string, userName string, password string)
}
