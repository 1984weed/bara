package contest

import (
	"bara/graphql_model"
	"bara/utils"
	"context"
	"fmt"
)

// Resolver represent the contest's resolver interface
type Resolver interface {
	GetContests(ctx context.Context, limit int, offset int) ([]*graphql_model.Contest, error)
	GetContest(ctx context.Context, slug string) (*graphql_model.Contest, error)
	CreateContest(ctx context.Context, contest graphql_model.NewContest) (*graphql_model.Contest, error)
	UpdateContest(ctx context.Context, contest graphql_model.NewContest) (*graphql_model.Contest, error)
	RemoveContest(ctx context.Context, slug string) error
}

type contestResolver struct {
	uc Usecase
}

// NewContestResolver initializes the contest/ resources graphql resolver
func NewContestResolver(uc Usecase) Resolver {
	return &contestResolver{uc}
}

// GetContests gets the list of contest
func (cr *contestResolver) GetContests(ctx context.Context, limit int, offset int) ([]*graphql_model.Contest, error) {
	contests, err := cr.uc.GetContests(limit, offset)

	if err != nil {
		return nil, err
	}

	resConttests := make([]*graphql_model.Contest, len(contests))
	for i, c := range contests {
		resConttests[i] = &graphql_model.Contest{
			ID:             fmt.Sprintln("%s", c.ID),
			ContestSlug:    c.Slug,
			Title:          c.Title,
			StartTimestamp: utils.GetISO8061(c.StartTime),
			Duration:       nil,
			Problems:       []*graphql_model.Problem{},
		}
	}
	return resConttests, nil
}

// GetContest gets a contest
func (cr *contestResolver) GetContest(ctx context.Context, slug string) (*graphql_model.Contest, error) {
	contest, err := cr.uc.GetContest(slug)

	if err != nil {
		return nil, err
	}

	problems := make([]*graphql_model.Problem, len(contest.Problems))

	for i, p := range contest.Problems {
		problems[i] = &graphql_model.Problem{
			ID:          int(p.ID),
			Slug:        p.Slug,
			Title:       p.Title,
			Description: p.Description,
		}
	}

	return &graphql_model.Contest{
		ID:             fmt.Sprintln("%s", contest.ID),
		ContestSlug:    contest.Slug,
		Title:          contest.Title,
		StartTimestamp: utils.GetISO8061(contest.StartTime),
		Duration:       nil,
		Problems:       problems,
	}, nil
}

func (cr *contestResolver) CreateContest(ctx context.Context, contest graphql_model.NewContest) (*graphql_model.Contest, error) {
	return nil, nil
}

func (cr *contestResolver) UpdateContest(ctx context.Context, contest graphql_model.NewContest) (*graphql_model.Contest, error) {
	return nil, nil
}

func (cr *contestResolver) RemoveContest(ctx context.Context, slug string) error {
	return nil
}
