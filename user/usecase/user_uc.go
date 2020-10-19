package usecase

import (
	"bara/model"
	"bara/user"
	"bara/user/domain"
	"bara/user/repository"
	"bara/utils"
	"context"
	"errors"
	"fmt"
	"time"
)

type userUsecase struct {
	runner         user.RepositoryRunner
	contextTimeout time.Duration
	userImage      repository.ImageUploader
}

// NewUserUsecase creates user usecase
func NewUserUsecase(runner user.RepositoryRunner, userImage repository.ImageUploader, contextTimeout time.Duration) user.Usecase {
	return &userUsecase{runner, contextTimeout, userImage}
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

func (u *userUsecase) GetUserByID(ctx context.Context, userID int64) (*model.Users, error) {
	user, err := u.runner.GetRepository().GetUserByID(ctx, userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByUserName ...
func (u *userUsecase) GetUserByUserName(ctx context.Context, userName string) (*model.Users, error) {
	user, err := u.runner.GetRepository().GetUserByUserName(ctx, userName)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser...
func (u *userUsecase) UpdateUser(ctx context.Context, userID int64, userForUpdate domain.UserForUpdate) (*model.Users, error) {
	if userForUpdate.Image != nil {
		data, err := utils.EncodeBase64Image(*userForUpdate.Image)

		if err != nil {
			return nil, err
		}

		uniqueID, err := u.userImage.UploadProfileImage(ctx, data)

		if err != nil {
			return nil, err
		}

		profileURL := u.userImage.GetProfileURL(uniqueID)
		userForUpdate.ImageURL = &profileURL
	}

	err := u.runner.RunInTransaction(func(r user.Repository) error {
		err := r.UpdateUser(ctx, userID, &userForUpdate)
		return err
	})

	if err != nil {
		return nil, err
	}

	return u.GetUserByID(ctx, userID)
}
