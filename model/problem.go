package model

import "time"

// Problems table data
type Problems struct {
	ID           int64
	Slug         string
	Title        string
	Description  string
	FunctionName string
	LanguageID   int64
	OutputType   string
	AuthorID     int64
	CreatedAt    time.Time `pg:"default:now()"`
	UpdatedAt    time.Time `pg:"default:now()"`
}

// ProblemsWithArgs represents Problems with many ProblemArgs
type ProblemsWithArgs struct {
	ID           int64
	Slug         string
	Title        string
	Description  string
	FunctionName string
	Args         []ProblemArgs
	OutputType   string
	AuthorID     int64
	CreatedAt    time.Time `pg:"default:now()"`
	UpdatedAt    time.Time `pg:"default:now()"`
}

// ProblemArgs table data
type ProblemArgs struct {
	ID        int64
	ProblemID int64
	OrderNo   int
	Name      string
	VarType   string
}
