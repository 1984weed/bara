package user

import "context"

type Usecase interface {
	Register(ctx context.Context, userName string, email string, password string) error
}
