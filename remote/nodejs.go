package remote

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"

	"github.com/go-pg/pg/v9"
)

type ConfigNode struct {
	ImageName string
}

type NodeJSClient struct {
	store *pg.DB
}

func NewNodeJsClient(store *pg.DB) *NodeJSClient {
	return &NodeJSClient{store: store}
}

type CodeResult struct {
	Status   string
	Result   string
	Expected string
	Time     int
}

type Question struct {
	ID           int64
	Slug         string
	Title        string
	Description  string
	FunctionName string
	ArgID        int
	AuthorID     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type QuestionArgs struct {
	ID         int64
	QuestionID int64
	OrderNo    int
	Name       string
	Type       string
}

type QuestionTestcases struct {
	ID         int64
	InputText  string
	OutputText string
}

func (n *NodeJSClient) Exec(slug string, typedCode string) (*CodeResult, string) {
	question := &Question{ID: 2}
	// args := &QuestionArgs{QuestionID: 2}
	err := n.store.Select(question)

	if err != nil {
		return nil, ""
	}
	// create args
	// err = n.store.Select(args)

	args := new([]QuestionArgs)
	err = n.store.Model(args).
		Where("question_args.question_id = ?", 2).
		Select()

	if err != nil {
		return nil, ""
	}

	qts := new([]QuestionTestcases)

	err = n.store.Model(qts).
		Where("question_testcases.question_id = ?", 2).
		Select()

	if err != nil {
		return nil, ""
	}

	testcase := ""
	inputCount := strings.Count((*qts)[0].InputText, "\\n") + 1

	testcase += fmt.Sprintln(len(*qts))
	testcase += fmt.Sprintln(inputCount)
	for _, qt := range *qts {
		inputCases := strings.Split(qt.InputText, "\\n")
		for _, in := range inputCases {
			testcase += fmt.Sprintln(in)
		}
		testcase += fmt.Sprintln(qt.OutputText)
	}

	execFile := fmt.Sprintf(nodeJsTemplate, typedCode, question.FunctionName)
	inputCommand := fmt.Sprintf(`echo -e %q > ./temp && echo -e %q | node temp`, execFile, testcase)

	fmt.Println(inputCommand)
	out, err := exec.Command("docker", "run", "node:12.10.0-alpine", "/bin/ash", "-c", inputCommand).Output()

	if err != nil {
		return nil, ""
	}

	bytesReader := bytes.NewReader(out)
	reader := bufio.NewReader(bytesReader)
	var result CodeResult
	stdoutArray := []string{}

	for {
		line, _, err := reader.ReadLine()

		stdoutArray = append(stdoutArray, fmt.Sprintf("%s", line))

		if err == io.EOF {
			break
		}

		err = json.Unmarshal([]byte(line), &result)
	}

	stdout := ""

	for i, s := range stdoutArray {
		if i < len(stdoutArray)-2 {
			stdout += s
		}
	}

	return &result, stdout
}
