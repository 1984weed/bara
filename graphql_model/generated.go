// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql_model

import (
	"fmt"
	"io"
	"strconv"
)

type CodeArg struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type CodeResult struct {
	Result *CodeResultDetail `json:"result"`
	Stdout string            `json:"stdout"`
}

type CodeResultDetail struct {
	Expected string  `json:"expected"`
	Input    *string `json:"input"`
	Result   string  `json:"result"`
	Status   string  `json:"status"`
	Time     int     `json:"time"`
}

type CodeSnippet struct {
	Code string       `json:"code"`
	Lang CodeLanguage `json:"lang"`
}

type NewQuestion struct {
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	FunctionName string       `json:"functionName"`
	OutputType   string       `json:"outputType"`
	LanguageID   CodeLanguage `json:"languageID"`
	ArgsNum      int          `json:"argsNum"`
	Args         []*CodeArg   `json:"args"`
	TestCases    []*TestCase  `json:"testCases"`
}

type Question struct {
	Slug         string         `json:"slug"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	CodeSnippets []*CodeSnippet `json:"codeSnippets"`
}

type SubmitCode struct {
	TypedCode string `json:"typedCode"`
	Lang      string `json:"lang"`
	Slug      string `json:"slug"`
}

type TestCase struct {
	Input  []*string `json:"input"`
	Output string    `json:"output"`
}

type CodeLanguage string

const (
	CodeLanguageJavaScript CodeLanguage = "JavaScript"
)

var AllCodeLanguage = []CodeLanguage{
	CodeLanguageJavaScript,
}

func (e CodeLanguage) IsValid() bool {
	switch e {
	case CodeLanguageJavaScript:
		return true
	}
	return false
}

func (e CodeLanguage) String() string {
	return string(e)
}

func (e *CodeLanguage) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CodeLanguage(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CodeLanguage", str)
	}
	return nil
}

func (e CodeLanguage) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
