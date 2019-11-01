package user

import (
	"bara/model"
	"context"
)

// Usecase ...
type Usecase interface {
	Register(ctx context.Context, userName string, email string, password string) (*model.Users, error)
	Login(ctx context.Context, userName string, email string, password string) (*model.Users, error)
}