package resolver

import (
	"bara/auth"
	"bara/graphql_model"
	"bara/user"
	"bara/user/domain"
	"bara/utils"
	"context"
	"strconv"
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
	var currentUser *auth.CurrentUser
	if currentUser = auth.ForContext(ctx); currentUser == nil {
		return nil, utils.GraphqlPermissionError()
	}

	user, err := u.uc.GetUserByID(ctx, currentUser.Sub)

	if err != nil {
		return nil, err
	}

	role := graphql_model.UserRoleNormal

	return &graphql_model.User{
		ID:          string(user.ID),
		DisplayName: user.DisplayName,
		UserName:    user.UserName,
		Email:       user.Email,
		Image:       user.ImageURL,
		Role:        &role,
		Bio:         user.Bio,
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
		ID:          strconv.Itoa(int(user.ID)),
		DisplayName: user.DisplayName,
		UserName:    user.UserName,
		Email:       user.Email,
		Image:       user.ImageURL,
		Bio:         user.Bio,
	}, nil
}

func (u *userResolver) UpdateMe(ctx context.Context, input graphql_model.UserInput) (*graphql_model.User, error) {
	var currentUser *auth.CurrentUser
	if currentUser = auth.ForContext(ctx); currentUser == nil {
		return nil, utils.GraphqlPermissionError()
	}
	userForUpdate := domain.UserForUpdate{
		UserName:    input.UserName,
		DisplayName: input.DisplayName,
		Email:       input.Email,
		Bio:         input.Bio,
		Image:       input.Image,
	}

	user, err := u.uc.UpdateUser(ctx, user.ID, userForUpdate)

	if err != nil {
		return nil, err
	}

	return &graphql_model.User{
		ID:          strconv.Itoa(int(user.ID)),
		DisplayName: user.DisplayName,
		UserName:    user.UserName,
		Email:       user.Email,
		Image:       user.ImageURL,
		Bio:         user.Bio,
	}, nil
}
