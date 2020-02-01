package contest

import (
	"bara/model"
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

// Repository represent the problem's store
type Repository interface {
	GetContests(ctx context.Context, limit, offset int) ([]model.Contests, error)
}

type contestRepositoryRunner struct {
	Conn *pg.DB
}

type RepositoryRunner interface {
	RunInTransaction(fn func(r Repository) error) error
	GetRepository() Repository
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

// NewProblemRepositoryRunner will create an object that represent the problem.Repository Runner Interface
func NewContestRepositoryRunner(Conn *pg.DB) RepositoryRunner {
	return &contestRepositoryRunner{Conn}
}

type contestRepository struct {
	Conn orm.DB
}

// newContestRepository will create an object that represent the problem.Repository interface
func newContestRepository(Conn orm.DB) Repository {
	return &contestRepository{Conn}
}
func (r *contestRepository) GetContests(ctx context.Context, limit, offset int) ([]model.Contests, error) {
	contests := new([]model.Contests)

	err := r.Conn.Model(contests).
		Select()

	return *contests, err
}
