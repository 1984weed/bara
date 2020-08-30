package resolver

import (
	"bara/auth"
	"bara/contest"
	"bara/graphql_model"
	"bara/model"
	"bara/problem"
	"bara/problem/domain"
	"bara/utils"
	"context"
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type problemResolver struct {
	uc problem.Usecase
	cc contest.Usecase
}

// NewProblemResolver initializes the problem/ resources graphql resolver
func NewProblemResolver(uc problem.Usecase, cc contest.Usecase) problem.Resolver {
	return &problemResolver{uc, cc}
}

var codeSlugToGraphQL = map[model.CodeLanguageSlug]graphql_model.CodeLanguage{
	model.JavaScript: graphql_model.CodeLanguageJavaScript,
}

var graphQlToCodeSlug = map[graphql_model.CodeLanguage]model.CodeLanguageSlug{
	graphql_model.CodeLanguageJavaScript: model.JavaScript,
}

// GetProblems get problem list
func (pr *problemResolver) GetProblems(ctx context.Context, limit int, offset int) ([]*graphql_model.Problem, error) {
	problems, err := pr.uc.GetProblems(ctx, limit, offset)

	if err != nil {
		return nil, err
	}

	graphProblems := make([]*graphql_model.Problem, len(problems))

	for i, p := range problems {
		graphProblems[i] = &graphql_model.Problem{
			ID:           int(p.ProblemID),
			Slug:         p.Slug,
			Title:        p.Title,
			Description:  p.Description,
			CodeSnippets: []*graphql_model.CodeSnippet{},
		}
	}

	return graphProblems, nil
}

// GetBySlug retrieves one problem by its slug
func (pr *problemResolver) GetBySlug(ctx context.Context, slug string) (*graphql_model.Problem, error) {
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
	codeArg := make([]*graphql_model.CodeArgType, len(p.ProblemArgs))

	for i, p := range p.ProblemArgs {
		codeArg[i] = &graphql_model.CodeArgType{
			Name: p.Name,
			Type: p.VarType,
		}
	}

	testcases := make([]*graphql_model.TestCaseType, len(p.ProblemTestcases))
	sampleTestCase := p.ProblemTestcases[0].Input

	for i, p := range p.ProblemTestcases {
		inputArray := p.ConvertInputArray()
		inputGraphqlArray := make([]*string, len(inputArray))

		for j := range inputArray {
			inputGraphqlArray[j] = &inputArray[j]
		}

		testcases[i] = &graphql_model.TestCaseType{
			Input:  inputGraphqlArray,
			Output: p.Output,
		}
	}

	return &graphql_model.Problem{
		ID:           int(p.ProblemID),
		Slug:         p.Slug,
		Title:        p.Title,
		Description:  p.Description,
		CodeSnippets: codeSnippets,
		ProblemDetailInfo: &graphql_model.ProblemDetailInfo{
			FunctionName: p.FunctionName,
			OutputType:   p.OutputType,
			ArgsNum:      len(p.ProblemArgs),
			Args:         codeArg,
			TestCases:    testcases,
		},
		SampleTestCase: &sampleTestCase,
	}, nil
}

// GetTestNewProblem does dry-run to create test new question
func (pr *problemResolver) GetTestNewProblem(ctx context.Context, input graphql_model.NewProblem) (*graphql_model.Problem, error) {
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
	return &graphql_model.Problem{
		Slug:         slug.Make(input.Title),
		Title:        input.Title,
		Description:  input.Description,
		CodeSnippets: codeSnippets,
	}, nil

}

func (pr *problemResolver) CreateProblem(ctx context.Context, input graphql_model.NewProblem) (*graphql_model.Problem, error) {
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
		Slug:         input.Slug,
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

	return &graphql_model.Problem{
		ID:    int(p.ProblemID),
		Title: p.Title,
		Slug:  p.Slug,
	}, nil
}

// UpdateProblem ...
func (pr *problemResolver) UpdateProblem(ctx context.Context, problemID int64, input graphql_model.NewProblem) (*graphql_model.Problem, error) {
	if user := auth.ForContext(ctx); user == nil {
		return nil, errors.New("Forbidden")
	}

	args := make([]domain.ProblemArgs, input.ArgsNum)
	for i, a := range input.Args {
		args[i] = domain.ProblemArgs{
			Name:    a.Name,
			VarType: a.Type,
		}
	}
	testcases := make([]domain.Testcase, input.TestCaseNum)
	for i := 0; i < input.TestCaseNum; i++ {
		t := input.TestCases[i]
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
		Slug:         input.Slug,
		Description:  input.Description,
		OutputType:   input.OutputType,
		FunctionName: input.FunctionName,
		ProblemArgs:  args,
		Testcases:    testcases,
	}
	p, err := pr.uc.UpdateProblem(ctx, problemID, problem)

	if err != nil {
		return nil, err
	}

	return &graphql_model.Problem{
		ID:    int(p.ProblemID),
		Title: p.Title,
		Slug:  p.Slug,
	}, nil
}

func (pr *problemResolver) SubmitProblem(ctx context.Context, input graphql_model.SubmitCode) (*graphql_model.CodeResult, error) {
	var user *auth.CurrentUser
	if user = auth.ForContext(ctx); user == nil {
		return nil, utils.GraphqlPermissionError()
	}

	domainCode := &domain.SubmitCode{
		LanguageSlug: graphQlToCodeSlug[graphql_model.CodeLanguage(input.Lang)],
		TypedCode:    input.TypedCode,
		ProblemSlug:  input.Slug,
	}
	result, err := pr.uc.SubmitProblem(ctx, domainCode, user.Sub)

	if err != nil {
		return nil, err
	}

	return &graphql_model.CodeResult{
		Result: &graphql_model.CodeResultDetail{
			Expected: result.Expected,
			Result:   string(result.Result),
			Status:   result.Status,
			Time:     result.Time,
			Input:    &result.Input,
		},
		Stdout: result.Output,
	}, nil
}

func (pr *problemResolver) SubmitContestCode(ctx context.Context, contestSlug string, input graphql_model.SubmitCode) (*graphql_model.CodeResult, error) {
	var user *auth.CurrentUser
	if user = auth.ForContext(ctx); user == nil {
		return nil, utils.GraphqlPermissionError()
	}

	domainCode := &domain.SubmitCode{
		LanguageSlug: graphQlToCodeSlug[graphql_model.CodeLanguage(input.Lang)],
		TypedCode:    input.TypedCode,
		ProblemSlug:  input.Slug,
	}
	result, err := pr.uc.SubmitProblem(ctx, domainCode, user.Sub)

	if err != nil {
		return nil, err
	}

	err = pr.cc.RegisterProblemResult(result, contestSlug, input.Slug, user.Sub)

	if err != nil {
		return nil, err
	}

	return &graphql_model.CodeResult{
		Result: &graphql_model.CodeResultDetail{
			Expected: result.Expected,
			Result:   string(result.Result),
			Status:   result.Status,
			Time:     result.Time,
			Input:    &result.Input,
		},
		Stdout: result.Output,
	}, nil
}

func (pr *problemResolver) TestRunCode(ctx context.Context, inputStr string, input graphql_model.SubmitCode) (*graphql_model.CodeResult, error) {
	domainCode := &domain.SubmitCode{
		LanguageSlug: graphQlToCodeSlug[graphql_model.CodeLanguage(input.Lang)],
		TypedCode:    input.TypedCode,
		ProblemSlug:  input.Slug,
	}

	result, err := pr.uc.RunProblem(ctx, domainCode, inputStr)

	if err != nil {
		return nil, err
	}

	return &graphql_model.CodeResult{
		Result: &graphql_model.CodeResultDetail{
			Expected: result.Expected,
			Result:   string(result.Result),
			Status:   result.Status,
			Time:     result.Time,
			Input:    &result.Input,
		},
		Stdout: result.Output,
	}, nil
}

func (pr *problemResolver) GetUsersSubmissionByProblemID(ctx context.Context, problemSlug string, limit, offset int) ([]*graphql_model.Submission, error) {
	var user *auth.CurrentUser
	if user = auth.ForContext(ctx); user == nil {
		return []*graphql_model.Submission{}, nil
	}

	submissions, err := pr.uc.GetUsersSubmissionByProblemID(ctx, user.ID, problemSlug, limit, offset)
	if err != nil {
		return []*graphql_model.Submission{}, err
	}

	results := make([]*graphql_model.Submission, len(submissions))

	for i, s := range submissions {
		results[i] = &graphql_model.Submission{
			ID:         fmt.Sprintln("%s", s.ID),
			LangSlug:   codeSlugToGraphQL[s.CodeLangSlug],
			RuntimeMs:  s.ExecTime,
			StatusSlug: s.StatusSlug,
			URL:        "undefined",
			Timestamp:  utils.GetISO8061(s.Timestamp),
		}
	}

	return results, nil
}
