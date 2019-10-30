package resolver

import (
	"bara/auth"
	"bara/graphql_model"
	"bara/model"
	"bara/problem"
	"bara/problem/domain"
	"context"
	"errors"

	"github.com/gosimple/slug"
)

type problemResolver struct {
	uc problem.Usecase
}

// NewProblemResolver initializes the problem/ resources graphql resolver
func NewProblemResolver(uc problem.Usecase) problem.Resolver {
	return &problemResolver{uc}
}

var codeSlugToGraphQL = map[model.CodeLanguageSlug]graphql_model.CodeLanguage{
	model.JavaScript: graphql_model.CodeLanguageJavaScript,
}

var graphQlToCodeSlug = map[graphql_model.CodeLanguage]model.CodeLanguageSlug{
	graphql_model.CodeLanguageJavaScript: model.JavaScript,
}

// GetBySlug retrieves one problem by its slug
func (pr *problemResolver) GetBySlug(ctx context.Context, slug string) (*graphql_model.Question, error) {
	if user := auth.ForContext(ctx); user == nil {
		return nil, errors.New("Forbidden")
	}
	p, err := pr.uc.GetBySlug(ctx, slug)

	if err != nil {
		return nil, err
	}

	codeSnippets := make([]*graphql_model.CodeSnippet, len(p.LanguageSlugs))

	for i, slug := range p.LanguageSlugs {
		codeSnippets[i] = &graphql_model.CodeSnippet{
			Code: p.MakeCodeSnippets()[i],
			Lang: codeSlugToGraphQL[slug],
		}
	}

	return &graphql_model.Question{
		Slug:         p.Slug,
		Title:        p.Title,
		Description:  p.Description,
		CodeSnippets: codeSnippets,
	}, nil
}

// GetTestNewProblem does dry-run to create test new question
func (pr *problemResolver) GetTestNewProblem(ctx context.Context, input graphql_model.NewQuestion) (*graphql_model.Question, error) {
	languages := []model.CodeLanguageSlug{model.JavaScript}

	codeSnippets := make([]*graphql_model.CodeSnippet, len(languages))
	args := make([]domain.ProblemArgs, input.ArgsNum)
	for i, a := range input.Args {
		args[i] = domain.ProblemArgs{
			Name:    a.Name,
			VarType: a.Type,
		}
	}
	problem := &domain.Problem{
		FunctionName:  input.FunctionName,
		ProblemArgs:   args,
		LanguageSlugs: languages,
		OutputType:    input.OutputType,
	}

	for i, slug := range []model.CodeLanguageSlug{model.JavaScript} {
		codeSnippets[i] = &graphql_model.CodeSnippet{
			Code: problem.MakeCodeSnippets()[i],
			Lang: codeSlugToGraphQL[slug],
		}
	}
	return &graphql_model.Question{
		Slug:         slug.Make(input.Title),
		Title:        input.Title,
		Description:  input.Description,
		CodeSnippets: codeSnippets,
	}, nil

}

func (pr *problemResolver) CreateProblem(ctx context.Context, input graphql_model.NewQuestion) (*graphql_model.Question, error) {
	args := make([]domain.ProblemArgs, input.ArgsNum)
	for i, a := range input.Args {
		args[i] = domain.ProblemArgs{
			Name:    a.Name,
			VarType: a.Type,
		}
	}
	testcases := make([]domain.Testcase, len(input.TestCases))
	for i, t := range input.TestCases {
		input := make([]string, len(t.Input))
		for i, in := range t.Input {
			input[i] = *in
		}
		testcases[i] = domain.Testcase{
			InputArray: input,
			Output:     t.Output,
		}
	}
	problem := &domain.NewProblem{
		Title:        input.Title,
		Description:  input.Description,
		OutputType:   input.OutputType,
		FunctionName: input.FunctionName,
		ProblemArgs:  args,
		Testcases:    testcases,
	}
	p, err := pr.uc.CreateProblem(ctx, problem)

	if err != nil {
		return nil, err
	}

	return &graphql_model.Question{
		Title: p.Title,
	}, nil
}

func (pr *problemResolver) SubmitProblem(ctx context.Context, input graphql_model.SubmitCode) (*graphql_model.CodeResult, error) {
	if user := auth.ForContext(ctx); user == nil {
		return nil, errors.New("Forbidden")
	}

	domainCode := &domain.SubmitCode{
		LanguageSlug: graphQlToCodeSlug[graphql_model.CodeLanguage(input.Lang)],
		TypedCode:    input.TypedCode,
		ProblemSlug:  input.Slug,
	}
	result, err := pr.uc.SubmitProblem(ctx, domainCode)

	if err != nil {
		return nil, err
	}

	return &graphql_model.CodeResult{
		Result: &graphql_model.CodeResultDetail{
			Expected: result.Expected,
			Result:   result.Result,
			Status:   result.Status,
			Time:     result.Time,
			Input:    &result.Input,
		},
		Stdout: result.Output,
	}, nil
}
