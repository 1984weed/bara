package resolver

import (
	"bara/graphql_model"
	"bara/model"
	"bara/problem"
	"context"
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

// GetBySlug(ctx context.Context, slug string) (*graphql_model.Question, error)
func (pr *problemResolver) GetBySlug(ctx context.Context, slug string) (*graphql_model.Question, error) {
	p, err := pr.uc.GetBySlug(ctx, slug)

	if err != nil {
		return nil, err
	}
	// question := new(remote.Question)

	// err := r.DB.Model(question).
	// 	Where("slug = ?", *slug).
	// 	Select()

	// if err != nil {
	// 	return nil, err
	// }

	// args := new([]remote.QuestionArgs)

	// err = r.DB.Model(args).
	// 	Where("question_args.question_id = ?", question.ID).
	// 	Select()
	// if err != nil {
	// 	return nil, err
	// }
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
