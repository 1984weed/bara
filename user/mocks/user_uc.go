package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// UserUsecase mock
type UserUsecase struct {
	mock.Mock
}

// Register mock
func (p *UserUsecase) Register(ctx context.Context, userName string, email string, password string) error {
	ret := p.Called(ctx, userName, email, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, userName, email, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
