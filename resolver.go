package bara

import (
	"bara/generated"
	"bara/graphql_model"
	"bara/problem"
	"bara/remote"
	"context"
	"fmt"
	"time"

	pg "github.com/go-pg/pg/v9"
	"github.com/gosimple/slug"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	DB               *pg.DB
	ProblemResolver  problem.Resolver
	WithoutContainer bool
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Questions(ctx context.Context, limit *int, offset *int) ([]*graphql_model.Question, error) {
	return []*graphql_model.Question{}, nil
}

func (r *queryResolver) Question(ctx context.Context, slug *string) (*graphql_model.Question, error) {
	return r.ProblemResolver.GetBySlug(ctx, *slug)
}

func (r *queryResolver) TestNewQuestion(ctx context.Context, input graphql_model.NewQuestion) (*graphql_model.Question, error) {
	return r.ProblemResolver.GetTestNewProblem(ctx, input)
}

// Mutation ...
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) SubmitCode(ctx context.Context, input graphql_model.SubmitCode) (*graphql_model.CodeResult, error) {
	jsClient := remote.NewNodeJsClient(r.DB, r.WithoutContainer)

	question := new(remote.Question)

	err := r.DB.Model(question).
		Where("slug = ?", input.Slug).
		Select()

	if err != nil {
		return nil, err
	}

	result, stdout := jsClient.Exec(question.ID, question.FunctionName, input.TypedCode)

	if result == nil {
		return nil, nil
	}

	return &graphql_model.CodeResult{
		Result: &graphql_model.CodeResultDetail{
			Expected: result.Expected,
			Result:   result.Result,
			Status:   result.Status,
			Time:     result.Time,
			Input:    &result.Input,
		},
		Stdout: stdout,
	}, nil
}

func (r *mutationResolver) CreateQuestion(ctx context.Context, input graphql_model.NewQuestion) (*graphql_model.Question, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	language := new(remote.CodeLanguage)

	err = r.DB.Model(language).
		Where("slug = ?", input.LanguageID.String()).
		Select()

	if err != nil {
		return nil, err
	}

	question := &remote.Question{
		Slug:         slug.Make(input.Title),
		Title:        input.Title,
		Description:  input.Description,
		FunctionName: input.FunctionName,
		OutputType:   input.OutputType,
		LanguageID:   language.ID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err = tx.Insert(question)
	if err != nil {
		return nil, err
	}
	for i, arg := range input.Args {
		err = tx.Insert(&remote.QuestionArgs{
			QuestionID: question.ID,
			OrderNo:    i + 1,
			Name:       arg.Name,
			VarType:    arg.Type,
		})
		if err != nil {
			return nil, err
		}
	}

	for _, testcase := range input.TestCases {
		inputString := ""
		for i, input := range testcase.Input {
			if i == 0 {
				inputString += fmt.Sprintf("%s", *input)
			} else {
				inputString += fmt.Sprintf("%s\n", *input)
			}
		}
		err = tx.Insert(&remote.QuestionTestcases{
			QuestionID: question.ID,
			InputText:  inputString,
			OutputText: testcase.Output,
		})
	}
	if err != nil {
		return nil, err
	}
	tx.Commit()

	return &graphql_model.Question{
		Slug:        slug.Make(input.Title),
		Title:       question.Title,
		Description: question.Description,
	}, nil
}
