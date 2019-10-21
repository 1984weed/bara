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
	// jsClient := remote.NewNodeJsClient(r.DB)
	// question := &Question{ID: 1}
	// result := jsClient.Exec("", "function helloWorld(){ console.log('Hellow world') }")

	// return &SubmitCode{
	// 	TypedCode: "",
	// 	Lang:      "",
	// 	Slug:      "",
	// }, nil
	return []*graphql_model.Question{}, nil
}

func (r *queryResolver) Question(ctx context.Context, slug *string) (*graphql_model.Question, error) {
	return r.ProblemResolver.GetBySlug(ctx, *slug)
	// p, err := r.problemResolver.GetBySlug(ctx, *slug)

	// if err != nil {
	// 	return &Question{}, err
	// }
	// // question := new(remote.Question)

	// // err := r.DB.Model(question).
	// // 	Where("slug = ?", *slug).
	// // 	Select()

	// // if err != nil {
	// // 	return nil, err
	// // }

	// // args := new([]remote.QuestionArgs)

	// // err = r.DB.Model(args).
	// // 	Where("question_args.question_id = ?", question.ID).
	// // 	Select()
	// // if err != nil {
	// // 	return nil, err
	// // }

	// return &Question{
	// 	Slug:        p.Slug,
	// 	Title:       p.Title,
	// 	Description: p.Description,
	// 	CodeSnippets: []*CodeSnippet{
	// 		{
	// 			Code: p.MakeCodeSnippets(), //makeSnippets(question.FunctionName, args, question.OutputType),
	// 			Lang: p.LanguageSlug,
	// 		},
	// 	},
	// }, nil
}

func (r *queryResolver) TestNewQuestion(ctx context.Context, input graphql_model.NewQuestion) (*graphql_model.Question, error) {
	args := make([]remote.QuestionArgs, len(input.Args))
	for i, arg := range input.Args {
		args[i] = remote.QuestionArgs{
			Name:    arg.Name,
			VarType: arg.Type,
		}
	}
	return &graphql_model.Question{
		Slug:        slug.Make(input.Title),
		Title:       input.Title,
		Description: input.Description,
		CodeSnippets: []*graphql_model.CodeSnippet{
			{
				Code: makeSnippets(input.FunctionName, &args, "number"),
				Lang: graphql_model.CodeLanguageJavaScript,
			},
		},
	}, nil

}

func makeSnippets(functionName string, args *[]remote.QuestionArgs, outputType string) string {
	argsString := ""
	explainArgs := ""
	for i, a := range *args {
		separator := ", "
		if i == 0 {
			separator = ""
		}
		explainArgs += fmt.Sprintln(fmt.Sprintf("* @param {%s} %s", convertJSTypeFromType(a.VarType), a.Name))
		argsString += fmt.Sprintf("%s%s", separator, a.Name)
	}
	explainArgs += fmt.Sprintf("* @return {%s}", convertJSTypeFromType(outputType))

	return fmt.Sprintf(`/**
%s 
*/
function %s(%s) {
	
}`, explainArgs, functionName, argsString)
}

func convertJSTypeFromType(typeStr string) string {
	switch typeStr {
	case "int", "double":
		return "number"
	case "int[]", "double[]":
		return "number[]"
	case "int[][]", "double[][]":
		return "number[][]"
	case "string":
		return "string"
	case "string[]":
		return "string[]"
	case "string[][]":
		return "string[][]"
	}
	return ""
}

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
				inputString += fmt.Sprintf("%s", input)
			} else {
				inputString += fmt.Sprintf("%s\n", input)
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
