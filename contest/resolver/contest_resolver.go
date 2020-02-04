package contest

import (
	contest "bara/contest/usecase"
	"bara/graphql_model"
	"bara/utils"
	"context"
	"fmt"
)

// Resolver represent the contest's resolver interface
type Resolver interface {
	GetContests(ctx context.Context, limit int, offset int) ([]*graphql_model.Contest, error)
}

type contestResolver struct {
	uc contest.Usecase
}

// NewContestResolver initializes the contest/ resources graphql resolver
func NewContestResolver(uc contest.Usecase) Resolver {
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
	return []*graphql_model.Contest{}, nil
}
