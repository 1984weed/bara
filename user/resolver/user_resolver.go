package resolver

import (
	"bara/auth"
	"bara/graphql_model"
	"bara/user"
	"bara/user/domain"
	"bara/utils"
	"context"
	"fmt"
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

func (u *userResolver) GetMe(ctx context.Context) (*graphql_model.User, error) {
	var currentUser *auth.CurrentUser
	if currentUser = auth.ForContext(ctx); currentUser == nil {
		return nil, utils.PermissionError
	}

	user, err := u.uc.GetUserByID(ctx, currentUser.Sub)

	if err != nil {
		return nil, utils.InternalServerError
	}

	role := graphql_model.UserRoleNormal

	return &graphql_model.User{
		ID:          fmt.Sprintf("%d", user.ID),
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
		return nil, utils.PermissionError
	}
	userForUpdate := domain.UserForUpdate{
		UserName:    input.UserName,
		DisplayName: input.DisplayName,
		Email:       input.Email,
		Bio:         input.Bio,
		Image:       input.Image,
	}

	user, err := u.uc.UpdateUser(ctx, currentUser.Sub, userForUpdate)

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
