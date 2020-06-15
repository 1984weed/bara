package user

import (
	"bara/model"
	"bara/user/domain"
	"context"
)

// Repository ...
type Repository interface {
	Register(ctx context.Context, user *model.Users) (*model.Users, error)
	GetUserByID(ctx context.Context, userID int64) (*model.Users, error)
	GetUserByUserName(ctx context.Context, userName string) (*model.Users, error)
	GetUserByEmail(ctx context.Context, email string) (*model.Users, error)
	UpdateUser(ctx context.Context, userID int64, userForUpdate *domain.UserForUpdate) error
}
