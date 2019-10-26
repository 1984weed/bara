package user_uc

import (
	"bara/model"
	"bara/user"
	"context"
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

func (u *userUsecase) Register(ctx context.Context, userName string, email string, password string) error {
	repo := u.runner.GetRepository()

	user, err := repo.GetUserByEmail(ctx, email)

	if err != nil {
		return err
	}

	if user != nil {
		return fmt.Errorf("Email: %s is already exists", email)
	}

	user, err = repo.GetUserByUserName(ctx, userName)

	if err != nil {
		return err
	}

	if user != nil {
		return fmt.Errorf("UserName: %s is already exists", userName)
	}

	user = &model.Users{
		UserName:  userName,
		Password:  password,
		Email:     email,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}

	_, err = repo.Register(ctx, user)

	return err
}
