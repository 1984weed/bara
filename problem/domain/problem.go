package domain

import (
	"bara/model"
	"fmt"
)

// Problem represents the problem model
type Problem struct {
	Slug          string
	Title         string
	Description   string
	LanguageSlugs []model.CodeLanguageSlug
	FunctionName  string
	ProblemArgs   []ProblemArgs
	OutputType    string
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
