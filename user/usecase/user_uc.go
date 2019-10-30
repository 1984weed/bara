package user_uc

import (
	"bara/model"
	"bara/user"
	"bara/utils"
	"context"
	"errors"
	"fmt"
	"time"
)

type userUsecase struct {
	runner         user.RepositoryRunner
	contextTimeout time.Duration
}

// NewUserUsecase creates user usecase
func NewUserUsecase(runner user.RepositoryRunner, contextTimeout time.Duration) user.Usecase {
	return &userUsecase{runner, contextTimeout}
}

func (u *userUsecase) Register(ctx context.Context, userName string, email string, password string) (*model.Users, error) {
	repo := u.runner.GetRepository()

	user, err := repo.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, fmt.Errorf("Email: %s is already exists", email)
	}

	user, err = repo.GetUserByUserName(ctx, userName)

	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, fmt.Errorf("UserName: %s is already exists", userName)
	}

	hashedPass, err := utils.HashPassword(password)

	if user != nil {
		return nil, err
	}

	user = &model.Users{
		UserName:  userName,
		Password:  hashedPass,
		Email:     email,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}

	me, err := repo.Register(ctx, user)

	if err != nil {
		return nil, err
	}

	return me, err
}
func (u *userUsecase) Login(ctx context.Context, userName string, email string, password string) (*model.Users, error) {
	repo := u.runner.GetRepository()

	user, err := repo.GetUserByUserName(ctx, userName)

	if err != nil {
		return nil, err
	}

	if user != nil {
		if user.Password == password {
			return user, nil
		}
	}

	user, err = repo.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("Not found")
	}

	return user, nil
}
