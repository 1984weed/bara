package usecase

import (
	"bara/problem"
	"time"
	// "github.com/bxcodec/go-clean-arch/article"
	// "github.com/bxcodec/go-clean-arch/author"
)

type problemUsecase struct {
	articleRepo problem.Repository
	// authorRepo     author.Repository
	contextTimeout time.Duration
}
