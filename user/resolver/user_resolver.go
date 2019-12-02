package resolver

import (
	"bara/auth"
	"bara/graphql_model"
	"bara/model"
	"bara/user"
	"bara/user/domain"
	"context"
	"errors"
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

func (u *userResolver) GetMe(ctx context.Context) (*graphql_model.User, error) {
	var user *model.Users
	if user = auth.ForContext(ctx); user == nil {
		return nil, errors.New("Forbidden")
	}

	user, err := u.uc.GetUserByID(ctx, user.ID)

	if err != nil {
		return nil, err
	}

	role := graphql_model.UserRoleNormal

	if domain.IsAdmin(user.UserName) {
		role = graphql_model.UserRoleAdmin
	}

	return &graphql_model.User{
		ID:       string(user.ID),
		RealName: user.RealName,
		UserName: user.UserName,
		Email:    user.Email,
		Image:    user.Image,
		Role:     &role,
		Bio:      user.Bio,
	}, nil
}

func (u *userResolver) GetUser(ctx context.Context, userName string) (*graphql_model.User, error) {
	user, err := u.uc.GetUserByUserName(ctx, userName)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return &graphql_model.User{
		ID:       string(user.ID),
		RealName: user.RealName,
		UserName: user.UserName,
		Email:    user.Email,
		Image:    user.Image,
		Bio:      user.Bio,
	}, nil
}

func (u *userResolver) UpdateUser(ctx context.Context, input graphql_model.UserInput) (*graphql_model.User, error) {
	userForUpdate := domain.UserForUpdate{
		UserName: input.UserName,
		RealName: input.RealName,
		Email:    input.Email,
		Bio:      input.Bio,
		Image:    input.Image,
	}

	_, err := u.uc.UpdateUser(ctx, userForUpdate)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
