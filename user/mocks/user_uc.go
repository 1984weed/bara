package mocks

import (
	"bara/model"
	"bara/user/domain"
	"context"

	"github.com/stretchr/testify/mock"
)

// UserUsecase mock
type UserUsecase struct {
	mock.Mock
}

// Register mock
func (u *UserUsecase) Register(ctx context.Context, userName string, email string, password string) (*model.Users, error) {
	return nil, nil
}

// func (p *UserUsecase) Register(ctx context.Context, userName string, email string, password string) error {
// 	ret := p.Called(ctx, userName, email, password)

// 	var r0 error
// 	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
// 		r0 = rf(ctx, userName, email, password)
// 	} else {
// 		r0 = ret.Error(0)
// 	}

// 	return r0
// }
func (u *UserUsecase) Login(ctx context.Context, userName string, email string, password string) (*model.Users, error) {
	return nil, nil
}

func (u *UserUsecase) GetUserByID(ctx context.Context, userID int64) (*model.Users, error) {
	ret := u.Called(ctx, userID)

	var r0 *model.Users
	if rf, ok := ret.Get(0).(func(context.Context, int64) *model.Users); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Users)
		}
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (u *UserUsecase) GetUserByUserName(ctx context.Context, userName string) (*model.Users, error) {
	ret := u.Called(ctx, userName)

	var r0 *model.Users
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Users); ok {
		r0 = rf(ctx, userName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Users)
		}
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1

}

func (u *UserUsecase) UpdateUser(ctx context.Context, userID int64, user domain.UserForUpdate) (*model.Users, error) {
	return nil, nil
}
