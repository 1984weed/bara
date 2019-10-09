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

func (r *queryResolver) Question(ctx context.Context, slug *string) (*Question, error) {
	question := new(remote.Question)

	err := r.DB.Model(question).
		Where("slug = ?", *slug).
		Select()

	if err != nil {
		return nil, err
	}

	args := new([]remote.QuestionArgs)

	err = r.DB.Model(args).
		Where("question_args.question_id = ?", question.ID).
		Select()
	if err != nil {
		return nil, err
	}

	return &Question{
		Slug:        question.Slug,
		Title:       question.Title,
		Description: question.Description,
		CodeSnippets: []*CodeSnippet{
			{
				Code: makeSnippets(question.FunctionName, args),
				Lang: CodeLanguageJavaScript,
			},
		},
	}, nil

}

func (r *queryResolver) TestNewQuestion(ctx context.Context, input NewQuestion) (*Question, error) {
	args := make([]remote.QuestionArgs, len(input.Args))
	for i, arg := range input.Args {
		args[i] = remote.QuestionArgs{
			Name: arg.Name,
			Type: arg.Type.String(),
		}
	}
	return &Question{
		Slug:        slug.Make(input.Title),
		Title:       input.Title,
		Description: input.Description,
		CodeSnippets: []*CodeSnippet{
			{
				Code: makeSnippets(input.FunctionName, &args),
				Lang: CodeLanguageJavaScript,
			},
		},
	}, nil

}

func makeSnippets(functionName string, args *[]remote.QuestionArgs) string {
	argsString := ""
	for i, a := range *args {
		separator := ", "
		if i == 0 {
			separator = ""
		}
		argsString += fmt.Sprintf("%s%s", separator, a.Name)
	}
	return fmt.Sprintf(`/**
 */
function %s(%s) {
	
};
	`, functionName, argsString)

}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) SubmitCode(ctx context.Context, input SubmitCode) (*CodeResult, error) {
	jsClient := remote.NewNodeJsClient(r.DB)

	question := new(remote.Question)

	err := r.DB.Model(question).
		Where("slug = ?", input.Slug).
		Select()

	if err != nil {
		return nil, err
	}

	result, stdout := jsClient.Exec(question.ID, question.FunctionName, input.TypedCode)

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
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		})
		if err != nil {
			return nil, err
		}
	}

	for _, testcase := range input.TestCases {
		inputString := ""
		for i, input := range testcase.Input {
			if i == 0 {
				inputString += fmt.Sprintf("%s", input)
			} else {
				inputString += fmt.Sprintf("%s\n", input)
			}
		}
		err = tx.Insert(&remote.QuestionTestcases{
			QuestionID: question.ID,
			InputText:  inputString,
			OutputText: testcase.Output,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
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
