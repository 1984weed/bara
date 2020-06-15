package repository

import (
	"bara/model"
	"bara/user"
	"bara/user/domain"
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type userRepositoryRunner struct {
	Conn *pg.DB
}

func (p *userRepositoryRunner) RunInTransaction(fn func(r user.Repository) error) error {
	return p.Conn.RunInTransaction(func(tx *pg.Tx) error {
		pr := newProblemRepository(interface{}(tx).(orm.DB))
		return fn(pr)
	})
}

func (p *userRepositoryRunner) GetRepository() user.Repository {
	return newProblemRepository(interface{}(p.Conn).(orm.DB))
}

// NewUserRepositoryRunner will create an object that represent the problem.Repository Runner Interface
func NewUserRepositoryRunner(Conn *pg.DB) user.RepositoryRunner {
	return &userRepositoryRunner{Conn}
}

type userRepository struct {
	Conn orm.DB
}

// newProblemRepository will create an object that represent the problem.Repository interface
func newProblemRepository(Conn orm.DB) user.Repository {
	return &userRepository{Conn}
}

func (u *userRepository) GetUserByID(ctx context.Context, userID int64) (*model.Users, error) {
	var user = &model.Users{ID: userID}

	err := u.Conn.Select(user)

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.Users, error) {
	var user = new(model.Users)
	err := u.Conn.Model(user).
		Where("email = ?", email).
		Select()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (u *userRepository) GetUserByUserName(ctx context.Context, userName string) (*model.Users, error) {
	var user = new(model.Users)
	err := u.Conn.Model(user).
		Where("user_name = ?", userName).
		Select()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (u *userRepository) Register(ctx context.Context, user *model.Users) (*model.Users, error) {
	err := u.Conn.Insert(user)

	if err != nil {
		return nil, err
	}

	return u.GetUserByUserName(ctx, user.UserName)
}

func (u *userRepository) UpdateUser(ctx context.Context, userID int64, userForUpdate *domain.UserForUpdate) error {
	var user = &model.Users{
		ID: userID,
	}
	var updateFlag = false

	if userForUpdate.DisplayName != nil {
		user.DisplayName = *userForUpdate.DisplayName
		updateFlag = true
	}

	if userForUpdate.Bio != nil {
		user.Bio = *userForUpdate.Bio
		updateFlag = true
	}

	if userForUpdate.UserName != nil {
		user.UserName = *userForUpdate.UserName
		updateFlag = true
	}

	if userForUpdate.Email != nil {
		user.Email = *userForUpdate.Email
		updateFlag = true
	}

	if userForUpdate.Image != nil {
		user.Image = *userForUpdate.Image
		updateFlag = true
	}

	if !updateFlag {
		return nil
	}
	_, err := u.Conn.Model(user).Where("id = ?", userID).UpdateNotZero()

	if err != nil {
		return err
	}
	return nil
}
