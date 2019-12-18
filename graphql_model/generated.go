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

type CodeArgType struct {
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

type NewProblem struct {
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	FunctionName string      `json:"functionName"`
	OutputType   string      `json:"outputType"`
	ArgsNum      int         `json:"argsNum"`
	Args         []*CodeArg  `json:"args"`
	TestCases    []*TestCase `json:"testCases"`
}

type Problem struct {
	ID                int                `json:"id"`
	Slug              string             `json:"slug"`
	Title             string             `json:"title"`
	Description       string             `json:"description"`
	CodeSnippets      []*CodeSnippet     `json:"codeSnippets"`
	ProblemDetailInfo *ProblemDetailInfo `json:"problemDetailInfo"`
	SampleTestCase    *string            `json:"sampleTestCase"`
}

type ProblemDetailInfo struct {
	FunctionName string          `json:"functionName"`
	OutputType   string          `json:"outputType"`
	ArgsNum      int             `json:"argsNum"`
	Args         []*CodeArgType  `json:"args"`
	TestCases    []*TestCaseType `json:"testCases"`
}

type RunCode struct {
	TypedCode string `json:"typedCode"`
	Lang      string `json:"lang"`
	Slug      string `json:"slug"`
}

type Submission struct {
	ID         string       `json:"id"`
	LangSlug   CodeLanguage `json:"langSlug"`
	RuntimeMs  int          `json:"runtimeMS"`
	StatusSlug string       `json:"statusSlug"`
	URL        string       `json:"url"`
	Timestamp  string       `json:"timestamp"`
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

type TestCaseType struct {
	Input  []*string `json:"input"`
	Output string    `json:"output"`
}

type User struct {
	ID       string    `json:"id"`
	RealName string    `json:"realName"`
	UserName string    `json:"userName"`
	Email    string    `json:"email"`
	Image    string    `json:"image"`
	Role     *UserRole `json:"role"`
	Bio      string    `json:"bio"`
}

type UserInput struct {
	RealName *string `json:"realName"`
	UserName *string `json:"userName"`
	Email    *string `json:"email"`
	Image    *string `json:"image"`
	Bio      *string `json:"bio"`
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

type UserRole string

const (
	UserRoleAdmin  UserRole = "admin"
	UserRoleNormal UserRole = "normal"
)

var AllUserRole = []UserRole{
	UserRoleAdmin,
	UserRoleNormal,
}

func (e UserRole) IsValid() bool {
	switch e {
	case UserRoleAdmin, UserRoleNormal:
		return true
	}
	return false
}

func (e UserRole) String() string {
	return string(e)
}

func (e *UserRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserRole", str)
	}
	return nil
}

func (e UserRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
