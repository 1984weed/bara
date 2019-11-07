package mocks

import (
	"bara/model"
	"bara/problem"
	"context"

	"github.com/stretchr/testify/mock"
)

// RepositoryRunner mock
type RepositoryRunner struct {
	mock.Mock
}

// GetRepository ...
func (r *RepositoryRunner) GetRepository() problem.Repository {
	ret := r.Called()

	var r0 problem.Repository
	if rf, ok := ret.Get(0).(func() problem.Repository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(problem.Repository)
		}
	}

	return r0
}

// RunInTransaction ...
func (r *RepositoryRunner) RunInTransaction(fn func(r problem.Repository) error) error {
	ret := r.Called(fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(r problem.Repository) error) error); ok {
		r0 = rf(fn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(error)
		}
	}

	return r0
}

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

// GetProblems ...
func (p *ProblemRepository) GetProblems(ctx context.Context, limit, offset int) ([]model.Problems, error) {
	ret := p.Called(ctx, limit, offset)
	var r0 []model.Problems

	if rf, ok := ret.Get(0).(func(context.Context, int, int) []model.Problems); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Problems)
		}
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTestcaseByProblemID ...
func (p *ProblemRepository) GetTestcaseByProblemID(ctx context.Context, problemID int64) ([]model.ProblemTestcases, error) {
	ret := p.Called(ctx, problemID)
	var r0 []model.ProblemTestcases

	if rf, ok := ret.Get(0).(func(context.Context, int64) []model.ProblemTestcases); ok {
		r0 = rf(ctx, problemID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.ProblemTestcases)
		}
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, problemID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveProblem ...
func (p *ProblemRepository) SaveProblem(ctx context.Context, problem *model.Problems) error {
	ret := p.Called(ctx, problem)
	var r0 error

	if rf, ok := ret.Get(0).(func(context.Context, *model.Problems) error); ok {
		r0 = rf(ctx, problem)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(error)
		}
	}

	return r0
}

// SaveProblemArgs ...
func (p *ProblemRepository) SaveProblemArgs(ctx context.Context, args *model.ProblemArgs) error {
	ret := p.Called(ctx, args)
	var r0 error

	if rf, ok := ret.Get(0).(func(context.Context, *model.ProblemArgs) error); ok {
		r0 = rf(ctx, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(error)
		}
	}

	return r0
}

// SaveProblemTestcase ...
func (p *ProblemRepository) SaveProblemTestcase(ctx context.Context, testcase *model.ProblemTestcases) error {
	ret := p.Called(ctx, testcase)
	var r0 error

	if rf, ok := ret.Get(0).(func(context.Context, *model.ProblemTestcases) error); ok {
		r0 = rf(ctx, testcase)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(error)
		}
	}

	return r0
}
