package resolver

import (
	"bara/user"
	"context"
)

type userResolver struct {
	uc user.Usecase
}

// NewUserResolver creates user resolver
func NewUserResolver(uc user.Usecase) user.Resolver {
	return &userResolver{
		uc,
	}
}

func (u *userResolver) Register(ctx context.Context, userName string, email string, password string) error {
	_, err := u.uc.Register(ctx, userName, email, password)

	if err != nil {
		return err
	}

	return nil
}

func (u *userResolver) Login(ctx context.Context, email string, userName string, password string) error {
	return nil
}
