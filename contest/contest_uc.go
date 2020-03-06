package contest

import (
	"bara/model"
	"bara/problem/domain"
	"sort"
)

// Usecase represent the problem's usecases
type Usecase interface {
	GetContests(limit, offset int) ([]model.Contests, error)
	GetContest(contestSlug string) (*ContestWithProblem, error)
	CreateContest(newcontest *NewContest) (*ContestWithProblem, error)
	UpdateContest(id model.ContestID, contest *NewContest) (*ContestWithProblem, error)
	UpdateRankingContest(contestSlug string) error
	DeleteContest(slug string) error
	RegisterProblemResult(result *domain.CodeResult, contestSlug string, problemSlug string, userID int64) error
}

type contestUsecase struct {
	runner RepositoryRunner
}

// NewContestUsecase creates new a contestUsecase object of contest.Usecase interface
func NewContestUsecase(runner RepositoryRunner) Usecase {
	return &contestUsecase{runner}
}

// GetContests ...
func (c *contestUsecase) GetContests(limit, offset int) ([]model.Contests, error) {
	contests, err := c.runner.GetRepository().GetContests(limit, offset)

	if err != nil {
		return []model.Contests{}, err
	}

	return contests, err
}

// GetContest ...
func (c *contestUsecase) GetContest(contestSlug string) (*ContestWithProblem, error) {
	contest, err := c.runner.GetRepository().GetContest(contestSlug)

	if err != nil {
		return nil, err
	}

	problems, err := c.runner.GetRepository().GetContestProblems(contestSlug)

	return &ContestWithProblem{
		ID:       int64(contest.ID),
		Slug:     contest.Slug,
		Problems: problems,
	}, err
}

// CreateContest
func (c *contestUsecase) CreateContest(newcontest *NewContest) (*ContestWithProblem, error) {
	var createdContest *model.Contests
	err := c.runner.RunInTransaction(func(r Repository) error {
		cc, err := r.CreateContest(newcontest)
		createdContest = cc

		cc, err := r.CreateContest(newcontest)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return c.GetContest(createdContest.Slug)
}

// UpdateContest
func (c *contestUsecase) UpdateContest(id model.ContestID, contest *NewContest) (*ContestWithProblem, error) {
	var createdContest *model.Contests
	err := c.runner.RunInTransaction(func(r Repository) error {
		cc, err := r.UpdateContest(id, contest)
		createdContest = cc

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return c.GetContest(createdContest.Slug)
}

// CreateContest
func (c *contestUsecase) DeleteContest(slug string) error {
	err := c.runner.RunInTransaction(func(r Repository) error {
		err := r.DeleteContest(slug)

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// RegisterResult the result on a contest
func (c *contestUsecase) RegisterProblemResult(result *domain.CodeResult, contestSlug string, problemSlug string, userID int64) error {
	err := c.runner.RunInTransaction(func(r Repository) error {
		err := r.CreateSubmitResult(result, contestSlug, problemSlug, userID)

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *contestUsecase) UpdateRankingContest(contestSlug string) error {
	contest, err := c.runner.GetRepository().GetContest(contestSlug)
	problems, err := c.runner.GetRepository().GetContestProblems(contestSlug)

	err = c.runner.RunInTransaction(func(r Repository) error {
		var userResultTimeMap map[int64]int
		for _, p := range problems {
			contestResults, err := r.GetContestProblemResult(contestSlug, p.Slug)

			if err != nil {
				return err
			}
			for _, c := range contestResults {
				_, ok := userResultTimeMap[c.UserID]

				timeSpend := int(c.CreatedAt.Unix() - contest.StartTime.Unix())
				if ok {
					userResultTimeMap[c.UserID] += timeSpend
				} else {
					userResultTimeMap[c.UserID] = timeSpend
				}
			}
		}
		type kv struct {
			Key   int64
			Value int
		}

		pl := make([]kv, len(userResultTimeMap))
		i := 0
		for k, v := range userResultTimeMap {
			pl[i] = kv{k, v}
			i++
		}
		sort.Slice(pl, func(i, j int) bool {
			return pl[i].Value > pl[j].Value
		})

		rankings := make([]ContestRanking, len(pl))

		for i, v := range pl {
			rankings[i] = ContestRanking{
				UserID:    v.Key,
				ContestID: contest.ID,
				Ranking:   i,
			}
		}

		err := r.UpdateContestRanking(rankings)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
