package model

import "time"

// Problem table data
type Problem struct {
	ID           int64
	Slug         string
	Title        string
	Description  string
	FunctionName string
	OutputType   string
	LanguageID   int64
	AuthorID     int64
	CreatedAt    time.Time `pg:"default:now()"`
	UpdatedAt    time.Time `pg:"default:now()"`
}

// ProblemArgs table data
type ProblemArgs struct {
	ID         int64
	QuestionID int64
	OrderNo    int
	Name       string
	VarType    string
}
