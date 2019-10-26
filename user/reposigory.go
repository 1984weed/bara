package user

import (
	"bara/model"
	"context"
)

// Repository ...
type Repository interface {
	Register(ctx context.Context, user *model.Users) (*model.Users, error)
	GetUserByUserName(ctx context.Context, userName string) (*model.Users, error)
	GetUserByEmail(ctx context.Context, email string) (*model.Users, error)
}
