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

func (n *NodeJSClient) Exec(slug string, typedCode string) (string, string) {
	question := &Question{ID: 2}
	err := n.store.Select(question)

	if err != nil {
		fmt.Println(err)
		return "", ""
	}
	mainFn := fmt.Sprintf("console.log(%s())", question.FunctionName)
	inputCommand := fmt.Sprintf(`echo "%s" > ./temp && echo "%s" >> ./temp  && node temp`, typedCode, mainFn)

	out, err := exec.Command("docker", "run", "node:12.10.0-alpine", "/bin/ash", "-c", inputCommand).Output()
	bytesReader := bytes.NewReader(out)
	reader := bufio.NewReader(bytesReader)
	var result string
	var stdout string
	for {
		line, _, err := reader.ReadLine()
		stdout += fmt.Sprintf("%s\n", line)
		result = fmt.Sprintf("%s\n", line)

		if err == io.EOF {
			break
		}
	}

	if err != nil {
		return "", ""
	}

	return result, stdout
	// err = store.QueryRow("select id,content,author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)

}
