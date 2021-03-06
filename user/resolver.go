package user

import (
	"bara/graphql_model"
	"context"
)

type Resolver interface {
	GetMe(ctx context.Context) (*graphql_model.User, error)
	GetUser(ctx context.Context, username string) (*graphql_model.User, error)
	UpdateMe(ctx context.Context, input graphql_model.UserInput) (*graphql_model.User, error)
}
