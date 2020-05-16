package domain

import (
	"bara/model"
	"fmt"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

// Problem represents the problem model
type Problem struct {
	ProblemID        int64
	Slug             string
	Title            string
	Description      string
	LanguageSlugs    []model.CodeLanguageSlug
	FunctionName     string
	ProblemArgs      []ProblemArgs
	ProblemTestcases []Testcase
	OutputType       string
}

func ConvertProblemFromTableModel(p model.Problems) *Problem {
	return &Problem{
		ProblemID:    p.ID,
		Slug:         p.Slug,
		Title:        p.Title,
		Description:  p.Description,
		FunctionName: p.FunctionName,
		OutputType:   p.OutputType,
	}
}

type ProblemArgs struct {
	Name    string
	VarType string
}

var mapDefaultCodeSnippet = map[model.CodeLanguageSlug]func(functionName string, args []ProblemArgs, outputType string) string{
	model.JavaScript: makeJSCodeSnippets,
}

func (p *Problem) MakeCodeSnippets() []string {
	snippets := make([]string, len(p.LanguageSlugs))
	for i, slug := range p.LanguageSlugs {
		snippets[i] = mapDefaultCodeSnippet[slug](p.FunctionName, p.ProblemArgs, p.OutputType)

	}
	return snippets
}

func makeJSCodeSnippets(functionName string, args []ProblemArgs, outputType string) string {
	argsString := ""
	explainArgs := ""
	for i, a := range args {
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

// NewProblem represents a new problem
type NewProblem struct {
	Title        string
	Slug         *string
	Description  string
	OutputType   string
	FunctionName string
	ProblemArgs  []ProblemArgs
	Testcases    []Testcase
}

// GetSlug returns slug from problem title
func (np *NewProblem) GetSlug() string {
	return slug.Make(np.Title)
}

// Testcase ...
type Testcase struct {
	InputArray []string
	Input      string
	Output     string
}

// GetInput returns all inputs with \n
func (t *Testcase) GetInput() string {
	inputString := ""

	for i, input := range t.InputArray {
		if i == len(t.InputArray)-1 {
			inputString += fmt.Sprintf("%s", input)
		} else {
			inputString += fmt.Sprintf("%s\n", input)
		}
	}

	return inputString
}

func (t *Testcase) ConvertInputArray() []string {
	inputCases := strings.Split(t.Input, "\n")

	return inputCases
}

func CreateTestcases(testcases []Testcase) []string {
	testcaseStrs := make([]string, len(testcases))
	inputCount := strings.Count(testcases[0].Input, "\n") + 1

	for i, qt := range testcases {
		testcase := fmt.Sprintln(inputCount)
		inputCases := strings.Split(qt.Input, "\n")
		for _, in := range inputCases {
			testcase += fmt.Sprintln(in)
		}
		testcase += fmt.Sprintln(qt.Output)
		testcaseStrs[i] = testcase
	}

	return testcaseStrs
}

func CreateTestcase(testcase Testcase) string {
	inputCount := strings.Count(testcase.Input, "\n") + 1

	testcaseStr := fmt.Sprintln(inputCount)
	inputCases := strings.Split(testcase.Input, "\n")
	for _, in := range inputCases {
		testcaseStr += fmt.Sprintln(in)
	}
	testcaseStr += fmt.Sprintln(testcase.Output)

	return testcaseStr
}

type SubmitCode struct {
	LanguageSlug model.CodeLanguageSlug
	TypedCode    string
	ProblemSlug  string
}

type CodeResultStatus string

const (
	StatusSuccess CodeResultStatus = "Success"
	StatusFail    CodeResultStatus = "Fail"
)

type CodeResult struct {
	Status   string
	Result   string
	Input    string
	Expected string
	Time     int
	StdOut   string
	Output   string
}

type CodeSubmission struct {
	ID           int64
	StatusSlug   string
	CodeLangSlug model.CodeLanguageSlug
	ExecTime     int
	Timestamp    time.Time
}
