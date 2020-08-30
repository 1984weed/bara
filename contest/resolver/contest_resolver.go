package resolver

import (
	"bara/auth"
	"bara/contest"
	"bara/contest/domain"
	"bara/graphql_model"
	"bara/model"
	"bara/utils"
	"context"
	"fmt"
	"strconv"
)

type contestResolver struct {
	uc contest.Usecase
}

// NewContestResolver initializes the contest/ resources graphql resolver
func NewContestResolver(uc contest.Usecase) contest.Resolver {
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
			ID:             strconv.Itoa(int(c.ID)),
			Slug:           c.Slug,
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

	resContest := contestToGraphqlContest(contest)

	user := auth.ForContext(ctx)
	if user != nil {
		contestProblemsStatus, err := cr.uc.GetContestProblemResult(model.ContestID(contest.ID), user.Sub)
		if err != nil {
			return nil, err
		}

		for i, p := range contest.Problems {
			_, ok := contestProblemsStatus[p.ID]
			resContest.Problems[i].UserResult = &graphql_model.ContestProblemsUserResult{
				Done: ok,
			}
		}
	}

	return resContest, nil
}

func (cr *contestResolver) CreateContest(ctx context.Context, contest graphql_model.NewContest) (*graphql_model.Contest, error) {
	c, err := cr.uc.CreateContest(graphqlContestToNewContest(contest))

	if err != nil {
		return nil, err
	}

	return contestToGraphqlContest(c), nil
}

func (cr *contestResolver) UpdateContest(ctx context.Context, contestID string, contest graphql_model.NewContest) (*graphql_model.Contest, error) {
	num, err := strconv.Atoi(contestID)

	if err != nil {
		return nil, err
	}

	c, err := cr.uc.UpdateContest(model.ContestID(num), graphqlContestToNewContest(contest))

	if err != nil {
		return nil, err
	}

	return contestToGraphqlContest(c), nil
}

func (cr *contestResolver) DeleteContest(ctx context.Context, slug string) error {
	err := cr.uc.DeleteContest(slug)

	if err != nil {
		return err
	}

	return nil
}

func (cr *contestResolver) UpdateRankingContest(ctx context.Context, slug string) (*graphql_model.Ranking, error) {
	err := cr.uc.UpdateRankingContest(slug)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func contestToGraphqlContest(contest *domain.ContestWithProblem) *graphql_model.Contest {
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
		ID:             fmt.Sprintf("%d", contest.ID),
		Slug:           contest.Slug,
		Title:          contest.Title,
		StartTimestamp: utils.GetISO8061(contest.StartTime),
		Duration:       nil,
		Problems:       problems,
	}
}

func graphqlContestToNewContest(contest graphql_model.NewContest) *domain.NewContest {
	startTime, err := utils.GetTimeFromString(contest.StartTimestamp)

	if err != nil {
		return nil
	}
	problemIDs := make([]int64, len(contest.ProblemIDs))

	for i, p := range contest.ProblemIDs {
		num, err := strconv.Atoi(p)

		if err != nil {
			continue
		}

		problemIDs[i] = int64(num)
	}

	return &domain.NewContest{
		Title:      contest.Title,
		Slug:       contest.Slug,
		StartTime:  startTime,
		ProblemIDs: problemIDs,
	}
}
