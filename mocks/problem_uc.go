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

func (p *ProblemUsecase) GetProblems(ctx context.Context, limit, offset int) ([]domain.Problem, error) {
	return nil, nil
}

func (p *ProblemUsecase) CreateProblem(ctx context.Context, input *domain.NewProblem) (*domain.Problem, error) {
	return nil, nil
}

func (p *ProblemUsecase) UpdateProblem(ctx context.Context, problemID int64, input *domain.NewProblem) (*domain.Problem, error) {
	return nil, nil
}
func (p *ProblemUsecase) SubmitProblem(ctx context.Context, code *domain.SubmitCode, userID int64) (*domain.CodeResult, error) {
	return nil, nil
}
func (p *ProblemUsecase) RunProblem(ctx context.Context, code *domain.SubmitCode, inputStr string) (*domain.CodeResult, error) {
	return nil, nil
}
func (p *ProblemUsecase) GetUsersSubmissionByProblemID(ctx context.Context, userID int64, problemSlug string, limit, offset int) ([]domain.CodeSubmission, error) {
	return nil, nil
}
