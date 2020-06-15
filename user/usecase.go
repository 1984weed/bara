package user

import (
	"bara/model"
	"bara/user/domain"
	"context"
)

// Usecase ...
type Usecase interface {
	Register(ctx context.Context, userName string, email string, password string) (*model.Users, error)
	Login(ctx context.Context, userName string, email string, password string) (*model.Users, error)
	GetUserByID(ctx context.Context, userID int64) (*model.Users, error)
	GetUserByUserName(ctx context.Context, userName string) (*model.Users, error)
	UpdateUser(ctx context.Context, userID int64, user domain.UserForUpdate) (*model.Users, error)
}
