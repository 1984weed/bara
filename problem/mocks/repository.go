package mocks

import (
	"bara/model"
	"context"

	"github.com/stretchr/testify/mock"
)

// ProblemRepository mock
type ProblemRepository struct {
	mock.Mock
}

// GetBySlug mock
func (p *ProblemRepository) GetBySlug(ctx context.Context, slug string) (*model.ProblemsWithArgs, error) {
	ret := p.Called(ctx, slug)
	var r0 *model.ProblemsWithArgs
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.ProblemsWithArgs); ok {
		r0 = rf(ctx, slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ProblemsWithArgs)
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
