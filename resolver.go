package bara

import (
	"bara/remote"
	"context"
	"fmt"

	pg "github.com/go-pg/pg/v9"
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

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) SubmitCode(ctx context.Context, input SubmitCode) (*CodeResult, error) {
	jsClient := remote.NewNodeJsClient(r.DB)

	result, stdout := jsClient.Exec("", input.TypedCode)

	fmt.Println(result)

	return &CodeResult{
		Result: &CodeResultDetail{
			Expected: result.Expected,
			Result:   result.Result,
			Status:   result.Status,
			Time:     result.Time,
		},
		Stdout: stdout,
	}, nil
}

func (r *mutationResolver) CreateQuestion(ctx context.Context, input NewQuestion) (*Question, error) {
	remote.Question{
		Slug:         "hogehoge",
		Title:        input.Title,
		Description:  input.Description,
		FunctionName: input.FunctionName,

		// ID           int64
		// Slug         string
		// Title        string
		// Description  string
		// FunctionName string
		// ArgID        int
		// AuthorID     int
		// CreatedAt    time.Time
		// UpdatedAt    time.Time
	}
	err = r.DB.Insert(&User{
		Name:   "root",
		Emails: []string{"root1@root", "root2@root"},
	})
	if err != nil {
		panic(err)
	}
	return &Question{
		Slug:        "testtest",
		Title:       input.Title,
		Description: input.Description,
	}, nil
}
