package rest

import (
	"bara"
	"bara/auth"
	"bara/problem"
	"bara/problem/domain"
	"errors"
	"net/http"

	"github.com/docker/distribution/context"
	"github.com/go-chi/render"
	"github.com/google/martian/log"
)

type ProblemResponse struct {
	Problem *NewProblem `json:"problem,omitempty"`
}

// ProblemAPI represents receivers for problem's api endpoints
type ProblemAPI struct {
	uc problem.Usecase
}

func NewProblemRestApi(uc problem.Usecase) *ProblemAPI {
	return &ProblemAPI{uc}
}

func (rd *ProblemResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

type NewProblemRequest struct {
	Problem *NewProblem `json:"problem,omitempty"`
}

func (a *NewProblemRequest) Bind(r *http.Request) error {
	return nil
}

// NewProblem is used when users create a new problem
type NewProblem struct {
	Title        string     `json:"title,omitempty"`
	Slug         string     `json:"slug,omitempty"`
	Description  string     `json:"description,omitempty"`
	OutputType   string     `json:"outputType,omitempty"`
	FunctionName string     `json:"functionName,omitempty"`
	Args         []Arg      `json:"args,omitempty"`
	Testcases    []Testcase `json:"testcases,omitempty"`
}

func (p *NewProblem) DomainProblemArgs() []domain.ProblemArgs {
	res := make([]domain.ProblemArgs, len(p.Args))

	for i, a := range p.Args {
		res[i] = domain.ProblemArgs{
			Name:    a.Name,
			VarType: a.Type,
		}
	}

	return res
}

func (p *NewProblem) DomainTestcases() []domain.Testcase {
	res := make([]domain.Testcase, len(p.Testcases))

	for i, t := range p.Testcases {
		res[i] = domain.Testcase{
			InputArray: t.Inputs,
			Output:     t.Output,
		}
	}

	return res
}

// ConvertDomainProblem ...
func ConvertDomainProblem(dProblem *domain.Problem) *NewProblem {
	args := make([]Arg, len(dProblem.ProblemArgs))

	for i, a := range dProblem.ProblemArgs {
		args[i] = Arg{
			Name: a.Name,
			Type: a.VarType,
		}
	}

	testcases := make([]Testcase, len(dProblem.ProblemTestcases))

	for i, t := range dProblem.ProblemTestcases {
		testcases[i] = Testcase{
			Inputs: t.InputArray,
			Output: t.Output,
		}
	}

	return &NewProblem{
		Title:        dProblem.Title,
		Slug:         dProblem.Slug,
		Description:  dProblem.Description,
		OutputType:   dProblem.OutputType,
		FunctionName: dProblem.FunctionName,
		Args:         args,
		Testcases:    testcases,
	}
}

// Arg ...
type Arg struct {
	Order int    `json:"order,omitempty"`
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
}

// Testcase ...
type Testcase struct {
	Inputs []string `json:"inputs,omitempty"`
	Output string   `json:"output,omitempty"`
}

// CreateProblem creates a problem
func (p *ProblemAPI) CreateProblem(w http.ResponseWriter, r *http.Request) {
	user := auth.ForContext(r.Context())
	if user == nil || !user.IsAdmin() {
		render.Render(w, r, bara.ErrUnauthorizedRequest())
		return
	}
	data := &NewProblemRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, bara.ErrInvalidRequest(err))
		return
	}

	problem := data.Problem

	createdProblem, err := p.uc.CreateProblem(context.Background(), &domain.NewProblem{
		Title:        problem.Title,
		Slug:         &problem.Slug,
		Description:  problem.Description,
		OutputType:   problem.OutputType,
		FunctionName: problem.FunctionName,
		ProblemArgs:  problem.DomainProblemArgs(),
		Testcases:    problem.DomainTestcases(),
	}, user.Sub)

	if err != nil {
		log.Errorf("Rest Api receiver (CreateProblem) happens error %v", err)
		render.Render(w, r, bara.ErrInvalidRequest(errors.New("An internal server error ocurred")))
		return
	}

	render.Status(r, http.StatusCreated)

	render.Render(w, r, &ProblemResponse{Problem: ConvertDomainProblem(createdProblem)})
}
