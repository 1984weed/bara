package contest

import (
	"bara/model"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

// Repository represent the problem's store
type Repository interface {
	GetContests(limit, offset int) ([]model.Contests, error)
	GetContest(slug string) (*model.Contests, error)
	GetContestProblems(slug string) ([]model.Problems, error)
	CreateContest(newContest *NewContest) (*model.Contests, error)
	UpdateContest(contestID model.ContestID, contest *NewContest) (*model.Contests, error)
	DeleteContest(slug string) error
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
			FROM problems p, contest c, contest_problems cp
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

	return updateContest, err
}

// DeleteContest
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
