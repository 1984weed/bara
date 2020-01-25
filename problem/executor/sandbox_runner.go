package executor

import (
	"bara/problem/domain"
	"bara/utils"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

type SandBoxRunner struct {
	Folder        string
	SandboxFile   string
	Command       string
	SubmittedCode string
	File          string
	TestcaseFile  string
	Testcase      []string
	ExeCommand    string
	Timeout       time.Duration
}

// NewSandBoxRunner
func NewSandBoxRunner(path string, folder string, command string, file string, testcase []string, exeCommand string, submittedCode string, timeout time.Duration) *SandBoxRunner {
	return &SandBoxRunner{
		Folder:        fmt.Sprintf("%s/%s", path, folder),
		SandboxFile:   fmt.Sprintf("%s/sandbox-cli", path),
		Command:       command,
		SubmittedCode: submittedCode,
		ExeCommand:    exeCommand,
		File:          fmt.Sprintf("%s/%s/%s", path, folder, file),
		TestcaseFile:  fmt.Sprintf("%s/%s/testcase", path, folder),
		Testcase:      testcase,
		Timeout:       timeout,
	}
}

// Exec returns...
func (s *SandBoxRunner) Exec() (*domain.CodeResult, error) {
	err := s.prepare()
	if err != nil {
		return nil, err
	}
	return s.run()
}

func (s *SandBoxRunner) prepare() error {
	err := os.MkdirAll(s.Folder, os.ModePerm)

	if err != nil {
		return err
	}
	// for i, t := range s.Testcase {
	// 	testcaseFile := utils.NewFileUtils(fmt.Sprintf("%s-%d", s.TestcaseFile, i))
	// 	err = testcaseFile.WriteCreateFile(t)
	// }

	// if err != nil {
	// 	return err
	// }

	execFile := utils.NewFileUtils(s.File)
	execFile.WriteCreateFile(s.SubmittedCode)

	return nil
}

func (s *SandBoxRunner) run() (*domain.CodeResult, error) {
	sandboxCommand := s.SandboxFile

	defer os.RemoveAll(s.Folder)

	if _, err := os.Stat(s.SandboxFile); os.IsNotExist(err) {
		sandboxCommand = ""
		before := time.Now()

		for _, t := range s.Testcase[1:] {
			ctx, cancel := context.WithTimeout(context.Background(), s.Timeout*time.Second)
			defer cancel()

			cmd := exec.CommandContext(ctx, s.ExeCommand, "-c", fmt.Sprintf("%s %s %s", sandboxCommand, s.Command, s.File))

			cmd.Stdin = strings.NewReader(t)
			var out bytes.Buffer
			cmd.Stdout = &out

			err := cmd.Run()

			if err != nil {
				return nil, err
			}
			bytesReader := bytes.NewReader(out.Bytes())
			reader := bufio.NewReader(bytesReader)
			var result domain.CodeResult
			stdoutArray := []string{}
			for {
				line, err := reader.ReadString('\n')

				if err == io.EOF {
					break
				}

				stdoutArray = append(stdoutArray, fmt.Sprintf("%s", line))
			}
			stdout := ""

			for i, s := range stdoutArray {
				if i < len(stdoutArray)-1 {
					stdout += s
				}
			}

			err = json.Unmarshal([]byte(stdoutArray[len(stdoutArray)-1]), &result)
			if result.Status == "Fail" {
				after := time.Now()
				return &domain.CodeResult{
					Status: "Fail",
					Time:   int(after.Unix() - before.Unix()),
				}, nil
			}
		}

		after := time.Now()
		return &domain.CodeResult{
			Status: "Success",
			Time:   int(after.Unix() - before.Unix()),
		}, nil
	}

	log.Println(fmt.Sprintf("%s %s %s", sandboxCommand, s.Command, s.File))
	defer os.RemoveAll(s.Folder)

	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, s.ExeCommand, "-c", fmt.Sprintf("cat %s | %s %s %s", s.TestcaseFile, sandboxCommand, s.Command, s.File))
	cmd.Env = []string{
		"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
	}
	cmd.SysProcAttr = &syscall.SysProcAttr{}

	return &domain.CodeResult{}, nil
}
