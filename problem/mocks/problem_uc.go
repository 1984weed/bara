package mocks

import (
	"bara/problem/domain"
	"context"

	"github.com/stretchr/testify/mock"
)

// ProblemUsecase mock
type ProblemUsecase struct {
	mock.Mock
}

// GetBySlug mock
func (p *ProblemUsecase) GetBySlug(ctx context.Context, slug string) (*domain.Problem, error) {
	ret := p.Called(ctx, slug)

	var r0 *domain.Problem
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Problem); ok {
		r0 = rf(ctx, slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Problem)
		}
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
