package model

import "time"

// Problems table data
type Problems struct {
	ID           int64
	Slug         string
	Title        string
	Description  string
	FunctionName string
	OutputType   string
	AuthorID     int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// ProblemsWithArgs represents Problems with many ProblemArgs
type ProblemsWithArgs struct {
	ID           int64
	Slug         string
	Title        string
	Description  string
	FunctionName string
	Args         []ProblemArgs
	Testcases    []ProblemTestcases
	OutputType   string
	AuthorID     int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// ProblemArgs table data
type ProblemArgs struct {
	ID        int64
	ProblemID int64
	OrderNo   int
	Name      string
	VarType   string
}

// ProblemTestcases table data
type ProblemTestcases struct {
	ID         int64
	ProblemID  int64
	InputText  string
	OutputText string
}

// ProblemUserResults table data
type ProblemUserResults struct {
	ID            int64
	ProblemID     int64
	UserID        int64
	SubmittedCode string
	Status        string
	CodeLangID    int64
	ExecTime      int
	CreatedAt     time.Time
}

// ProblemUserSubmission for retreving the submission data
type ProblemUserSubmission struct {
	ID            int64
	SubmittedCode string
	Status        string
	CodeLangSlug  CodeLanguageSlug
	ExecTime      int
	CreatedAt     time.Time
}
