package usecase

import (
	"bara/contest"
	"bara/contest/domain"
	"bara/contest/repository"
	"bara/model"
	problem_domain "bara/problem/domain"
	"sort"
)

type contestUsecase struct {
	runner repository.RepositoryRunner
}

// NewContestUsecase creates new a contestUsecase object of contest.Usecase interface
func NewContestUsecase(runner repository.RepositoryRunner) contest.Usecase {
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
func (c *contestUsecase) GetContest(contestSlug string) (*domain.ContestWithProblem, error) {
	contest, err := c.runner.GetRepository().GetContest(contestSlug)

	if err != nil {
		return nil, err
	}

	problems, err := c.runner.GetRepository().GetContestProblems(contestSlug)

	return &domain.ContestWithProblem{
		ID:        int64(contest.ID),
		Title:     contest.Title,
		Slug:      contest.Slug,
		StartTime: contest.StartTime,
		Problems:  problems,
	}, err
}

// CreateContest
func (c *contestUsecase) CreateContest(newcontest *domain.NewContest) (*domain.ContestWithProblem, error) {
	var createdContest *model.Contests
	err := c.runner.RunInTransaction(func(r contest.Repository) error {
		cc, err := r.CreateContest(newcontest)
		createdContest = cc

		contestProblems := make([]domain.ContestProblemID, len(newcontest.ProblemIDs))

		for i, pID := range newcontest.ProblemIDs {
			contestProblems[i] = domain.ContestProblemID{
				ContestID: cc.ID,
				ProblemID: pID,
			}
		}

		err = r.RegisterContestProblem(contestProblems)

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
func (c *contestUsecase) UpdateContest(id model.ContestID, newContest *domain.NewContest) (*domain.ContestWithProblem, error) {
	var updatedContest *model.Contests
	err := c.runner.RunInTransaction(func(r contest.Repository) error {
		cc, err := r.UpdateContest(id, newContest)
		updatedContest = cc

		err = r.DeleteContestProblem(id)

		if err != nil {
			return err
		}

		if len(newContest.ProblemIDs) > 0 {
			contestProblems := make([]domain.ContestProblemID, len(newContest.ProblemIDs))

			for i, pID := range newContest.ProblemIDs {
				contestProblems[i] = domain.ContestProblemID{
					ContestID: cc.ID,
					ProblemID: pID,
				}
			}

			err = r.RegisterContestProblem(contestProblems)
			if err != nil {
				return err
			}

		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return c.GetContest(updatedContest.Slug)
}

// CreateContest
func (c *contestUsecase) DeleteContest(slug string) error {
	err := c.runner.RunInTransaction(func(r contest.Repository) error {
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
func (c *contestUsecase) RegisterProblemResult(result *problem_domain.CodeResult, contestSlug string, problemSlug string, userID int64) error {
	err := c.runner.RunInTransaction(func(r contest.Repository) error {
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
	targetContest, err := c.runner.GetRepository().GetContest(contestSlug)
	problems, err := c.runner.GetRepository().GetContestProblems(contestSlug)

	err = c.runner.RunInTransaction(func(r contest.Repository) error {
		var userResultTimeMap map[int64]int
		for _, p := range problems {
			contestResults, err := r.GetContestProblemResult(contestSlug, p.Slug)

			if err != nil {
				return err
			}
			for _, c := range contestResults {
				_, ok := userResultTimeMap[c.UserID]

				timeSpend := int(c.CreatedAt.Unix() - targetContest.StartTime.Unix())
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

		rankings := make([]domain.ContestRanking, len(pl))

		for i, v := range pl {
			rankings[i] = domain.ContestRanking{
				UserID:    v.Key,
				ContestID: targetContest.ID,
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

// GetContestProblemResult ...
func (c *contestUsecase) GetContestProblemResult(contestID model.ContestID, userID int64) (map[int64]domain.UserContestProblemResult, error) {
	result, err := c.runner.GetRepository().GetContestProblemsUserResults(contestID, userID)

	if err != nil {
		return nil, err
	}

	res := map[int64]domain.UserContestProblemResult{}

	for _, r := range result {
		rp, ok := res[r.ProblemID]

		if ok {
			if rp.Done {
				continue
			}
		}

		res[r.ProblemID] = domain.UserContestProblemResult{
			ContestID: contestID,
			ProblemID: r.ProblemID,
			Done:      r.Status == "success",
		}
	}

	return res, nil
}
