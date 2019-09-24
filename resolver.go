package bara

import (
	"bara/remote"
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	DB *pg.DB
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Questions(ctx context.Context, limit *int, offset *int) ([]*Question, error) {
	// jsClient := remote.NewNodeJsClient(r.DB)
	// question := &Question{ID: 1}
	// result := jsClient.Exec("", "function helloWorld(){ console.log('Hellow world') }")

	// return &SubmitCode{
	// 	TypedCode: "",
	// 	Lang:      "",
	// 	Slug:      "",
	// }, nil
	return []*Question{}, nil
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// type MutationResolver interface {
// 	SubmitCode(ctx context.Context, input SubmitCode) (*CodeResult, error)
// }

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) SubmitCode(ctx context.Context, input SubmitCode) (*CodeResult, error) {
	jsClient := remote.NewNodeJsClient(r.DB)
	// question := &Question{ID: 1}
	result, stdout := jsClient.Exec("", "function helloWorld(){ console.log('Hellow world') }")

	fmt.Println(result)

	return &CodeResult{
		Result: result,
		Stdout: stdout,
	}, nil
}
