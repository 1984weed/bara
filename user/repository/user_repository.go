package repository

import (
	"bara/model"
	"bara/user"
	"context"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
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

func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.Users, error) {
	var user = new(model.Users)
	err := u.Conn.Model(user).
		Where("email = ?", email).
		Select()

	if err != nil {
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
		return nil, err
	}

	return user, nil
}

func (u *userRepository) Resister(ctx context.Context, user *model.Users) error {
	return u.Conn.Insert(user)
}
