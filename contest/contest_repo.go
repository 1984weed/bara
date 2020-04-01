package contest

import (
	"bara/model"
	"bara/problem/domain"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

// Repository represent the problem's store
type Repository interface {
	GetContests(limit, offset int) ([]model.Contests, error)
	GetContest(slug string) (*model.Contests, error)
	GetContestProblems(slug string) ([]model.Problems, error)
	GetContestProblemResult(contestSlug string, problemSlug string) ([]model.ContestUserProblemSuccess, error)
	UpdateContestRanking(ranking []ContestRanking) error
	CreateContest(newContest *NewContest) (*model.Contests, error)
	RegisterContestProblem(contestProblems []ContestProblemID) error
	DeleteContestProblem(contestID model.ContestID) error
	UpdateContest(contestID model.ContestID, contest *NewContest) (*model.Contests, error)
	DeleteContest(slug string) error
	CreateSubmitResult(result *domain.CodeResult, contestSlug string, problemSlug string, userID int64) error
}

// RepositoryRunner can run repo
type RepositoryRunner interface {
	RunInTransaction(fn func(r Repository) error) error

	GetRepository() Repository
}

// NewProblemRepositoryRunner will create an object that represent the problem.Repository Runner Interface
func NewContestRepositoryRunner(Conn *pg.DB) RepositoryRunner {
	return &contestRepositoryRunner{Conn}
}

type contestRepositoryRunner struct {
	Conn *pg.DB
}

func (p *contestRepositoryRunner) RunInTransaction(fn func(r Repository) error) error {
	return p.Conn.RunInTransaction(func(tx *pg.Tx) error {
		pr := newContestRepository(interface{}(tx).(orm.DB))
		return fn(pr)
	})
}

func (p *contestRepositoryRunner) GetRepository() Repository {
	return newContestRepository(interface{}(p.Conn).(orm.DB))
}

type contestRepository struct {
	Conn orm.DB
}

// newContestRepository will create an object that represent the problem.Repository interface
func newContestRepository(Conn orm.DB) Repository {
	return &contestRepository{Conn}
}

func (r *contestRepository) GetContests(limit, offset int) ([]model.Contests, error) {
	contests := new([]model.Contests)

	err := r.Conn.Model(contests).
		Limit(limit).
		Offset(offset).
		Select()

	return *contests, err
}

func (r *contestRepository) GetContest(slug string) (*model.Contests, error) {
	contest := new(model.Contests)
	err := r.Conn.Model(contest).
		Where("contests.slug = ?", slug).
		Select()

	if err != nil {
		return nil, err
	}

	return contest, err
}

func (r *contestRepository) GetContestProblems(slug string) ([]model.Problems, error) {
	var problems []model.Problems

	_, err := r.Conn.Query(
		&problems, `
				SELECT 
					p.*
				FROM problems p, contests c, contest_problems cp
				WHERE c.slug = ?
				AND cp.problem_id = p.id 
				AND c.id = cp.contest_id
				ORDER BY cp.order_id DESC
			`, slug)

	if err != nil {
		return []model.Problems{}, err
	}

	return problems, err

}
func (r *contestRepository) UpdateContestRanking(rankings []ContestRanking) error {
	res := make([]model.ContestUserResults, len(rankings))

	for i, r := range rankings {
		res[i] = model.ContestUserResults{
			ContestID: r.ContestID,
			UserID:    r.UserID,
			Ranking:   r.Ranking,
		}
	}

	return r.Conn.Insert(res)
}

func (r *contestRepository) GetContestProblemResult(contestSlug string, problemSlug string) ([]model.ContestUserProblemSuccess, error) {
	var res []model.ContestUserProblemSuccess

	_, err := r.Conn.Query(
		&res, `
			SELECT cpur.user_id, cpur.created_at, cpur.exec_time
			FROM problems p, contests c , contest_problem_user_results cpur
			INNER JOIN (
			SELECT min(created_at) AS min_date FROM contest_problem_user_results WHERE status = 'success' GROUP BY user_id ) cm ON cm.min_date = created_at
			WHERE 
			p.id = cpur.problem_id
			AND c.id = cpur.contest_id
			AND c.slug = ?
			AND p.slug = ?
			AND cpur.status = 'success' 
			GROUP BY cpur.user_id HAVING cpur.created_at = MIN(cpur.created_at)
			ORDER BY cpur.created_at
			`, contestSlug, problemSlug)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// CreateContest
func (r *contestRepository) CreateContest(newContest *NewContest) (*model.Contests, error) {
	contest := &model.Contests{
		Slug:      newContest.Slug,
		Title:     newContest.Title,
		StartTime: newContest.StartTime,
	}
	err := r.Conn.Insert(contest)

	if err != nil {
		return nil, err
	}

	return contest, err
}

// RegisterContestProblem...
func (r *contestRepository) RegisterContestProblem(contestProblems []ContestProblemID) error {
	contestsPs := make([]model.ContestProblems, len(contestProblems))

	for i, cp := range contestProblems {
		contestsPs[i] = model.ContestProblems{
			ContestID: cp.ContestID,
			ProblemID: cp.ProblemID,
			OrderID:   i,
		}
	}
	err := r.Conn.Insert(&contestsPs)

	if err != nil {
		return err
	}

	return nil
}

// DeleteContestProblem...
func (r *contestRepository) DeleteContestProblem(contestID model.ContestID) error {
	deleteContestProblems := &model.ContestProblems{
		ContestID: contestID,
	}

	_, err := r.Conn.Model(deleteContestProblems).Where("contest_id = ?contest_id").Delete()

	if err != nil {
		return err
	}
	return nil
}

// UpdateContest
func (r *contestRepository) UpdateContest(contestID model.ContestID, contest *NewContest) (*model.Contests, error) {
	updateContest := &model.Contests{
		ID:        contestID,
		Slug:      contest.Slug,
		Title:     contest.Title,
		StartTime: contest.StartTime,
	}
	err := r.Conn.Update(updateContest)

	if err != nil {
		return nil, err
	}

	return updateContest, nil
}

// DeleteContest is ...
func (r *contestRepository) DeleteContest(slug string) error {
	deleteContest := &model.Contests{
		Slug: slug,
	}
	err := r.Conn.Delete(deleteContest)

	if err != nil {
		return err
	}

	return nil
}

// CreateSubmitResult is ...
func (r *contestRepository) CreateSubmitResult(result *domain.CodeResult, contestSlug string, problemSlug string, userID int64) error {
	contest := new(model.Contests)

	err := r.Conn.Model(contest).
		Where("contests.slug = ?", contestSlug).
		Select()

	if err != nil {
		return err
	}

	problem := new(model.Problems)

	err = r.Conn.Model(problem).
		Where("problems.slug = ?", problemSlug).
		Select()

	if err != nil {
		return err
	}

	problemResult := &model.ContestProblemUserResults{
		ContestID: contest.ID,
		ProblemID: problem.ID,
		UserID:    userID,
		Status:    result.Status,
		ExecTime:  result.Time,
	}

	err = r.Conn.Insert(problemResult)

	if err != nil {
		return err
	}

	return nil
}
