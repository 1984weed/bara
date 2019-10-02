package bara

import (
	"bara/remote"
	"context"
	"fmt"
	"time"

	pg "github.com/go-pg/pg/v9"
	"github.com/gosimple/slug"
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
	tx, err := r.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// tx, err := db.Begin()
	// if err != nil {
	//     return err
	// }
	// // Rollback tx on error.
	// defer tx.Rollback()

	// var counter int
	// _, err = tx.QueryOne(
	//     pg.Scan(&counter), `SELECT counter FROM tx_test FOR UPDATE`)
	// if err != nil {
	//     return err
	// }

	// counter++

	// _, err = tx.Exec(`UPDATE tx_test SET counter = ?`, counter)
	// if err != nil {
	//     return err
	// }

	// return tx.Commit()

	question := &remote.Question{
		Slug:         slug.Make(input.Title),
		Title:        input.Title,
		Description:  input.Description,
		FunctionName: input.FunctionName,
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
			Type:       convertArgsType(arg.Type),
		})
		if err != nil {
			return nil, err
		}
	}

	for _, testcase := range input.TestCases {
		inputString := ""
		for _, input := range testcase.Input {
			inputString += fmt.Sprintf("%s\n", input)
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

	return &Question{
		Slug:        "testtest",
		Title:       question.Title,
		Description: question.Description,
	}, nil
}

func convertArgsType(argType TestCaseArgType) string {
	switch argType {
	case TestCaseArgTypeNumber:
		return "num"
	case TestCaseArgTypeString:
		return "string"
	}
	return "num"
}
