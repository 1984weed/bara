package user

import (
	"bara/graphql_model"
	"context"
)

type Resolver interface {
	GetMe(ctx context.Context) (*graphql_model.User, error)
	GetUser(ctx context.Context, username string) (*graphql_model.User, error)
}
