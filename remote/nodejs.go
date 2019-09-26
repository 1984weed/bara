package remote

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
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

func (n *NodeJSClient) Exec(slug string, typedCode string) (string, string) {
	question := &Question{ID: 2}
	// args := &QuestionArgs{QuestionID: 2}
	err := n.store.Select(question)

	if err != nil {
		fmt.Println(err)
		return "", ""
	}
	// create args
	// err = n.store.Select(args)

	args := new([]QuestionArgs)
	err = n.store.Model(args).
		Where("question_args.question_id = ?", 2).
		Select()

	if err != nil {
		return "", ""
	}

	// main function
	mainFn := fmt.Sprintf("console.log(%s())", question.FunctionName)
	inputCommand := fmt.Sprintf(`time echo "%s" > ./temp && echo "%s" >> ./temp  && node temp`, typedCode, mainFn)

	out, err := exec.Command("docker", "run", "node:12.10.0-alpine", "/bin/ash", "-c", inputCommand).Output()

	if err != nil {
		return "", ""
	}

	bytesReader := bytes.NewReader(out)
	reader := bufio.NewReader(bytesReader)
	var result string
	stdoutArray := []string{}

	for {
		line, _, err := reader.ReadLine()
		fmt.Println(string(line))

		stdoutArray = append(stdoutArray, fmt.Sprintf("%s", line))

		if err == io.EOF {
			break
		}

		result = fmt.Sprintf("%s", line)
	}

	stdout := ""

	for i, s := range stdoutArray {
		if i < len(stdoutArray)-2 {
			stdout += s
		}
	}

	return result, stdout
}
