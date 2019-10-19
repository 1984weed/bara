package remote

import (
	"bara/utils"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/go-pg/pg/v9"
)

type ConfigNode struct {
	ImageName string
}

type NodeJSClient struct {
	store            *pg.DB
	withoutContainer bool
}

func NewNodeJsClient(store *pg.DB, withoutContainer bool) *NodeJSClient {
	return &NodeJSClient{store: store, withoutContainer: withoutContainer}
}

type CodeResult struct {
	Status   string
	Result   string
	Input    string
	Expected string
	Time     int
}

type Question struct {
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

type CodeLanguage struct {
	ID   int64
	Name string
	Slug string
}

type QuestionArgs struct {
	ID         int64
	QuestionID int64
	OrderNo    int
	Name       string
	VarType    string
}

type QuestionTestcases struct {
	ID         int64
	QuestionID int64
	InputText  string `pg:",notnull"`
	OutputText string
}

func (n *NodeJSClient) Exec(questionID int64, functionName string, typedCode string) (*CodeResult, string) {
	args := new([]QuestionArgs)
	err := n.store.Model(args).
		Where("question_args.question_id = ?", questionID).
		Select()

	if err != nil {
		return nil, ""
	}

	qts := new([]QuestionTestcases)

	err = n.store.Model(qts).
		Where("question_testcases.question_id = ?", questionID).
		Select()

	if err != nil {
		return nil, ""
	}

	testcase := ""
	inputCount := strings.Count((*qts)[0].InputText, "\n") + 1

	testcase += fmt.Sprintln(len(*qts))
	testcase += fmt.Sprintln(inputCount)
	for _, qt := range *qts {
		inputCases := strings.Split(qt.InputText, "\n")
		for _, in := range inputCases {
			testcase += fmt.Sprintln(in)
		}
		testcase += fmt.Sprintln(qt.OutputText)
	}

	dir, err := os.Getwd()
	if err != nil {
		return nil, ""
	}

	codeCompile := CodeCompileMap[JavaScript]
	var machine MachineType
	if n.withoutContainer {
		machine = LocalMac
	} else {
		machine = Container
	}
	machineType := MachineExecMap[machine]
	execStr := fmt.Sprintf(codeCompile.PrepareCode, typedCode, functionName)

	sandbox := NewSandBoxRunner(dir, fmt.Sprintf(`folder-%s`, utils.RandomString(10)), codeCompile.Command, codeCompile.FileName, testcase, machineType, execStr, 600, JavaScript)
	out, err := sandbox.Exec()

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
