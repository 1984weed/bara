package user

type RepositoryRunner interface {
	RunInTransaction(fn func(r Repository) error) error
	GetRepository() Repository
}
