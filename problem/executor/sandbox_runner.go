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
	"os"
	"os/exec"
	"strings"
	"time"

	seccomp "github.com/elastic/go-seccomp-bpf"
)

var (
	policyFile string
	noNewPrivs bool
)

type SandBoxRunner struct {
	Folder        string
	SandboxFile   string
	Command       string
	SubmittedCode string
	File          string
	TestcaseFile  string
	Testcase      []domain.Testcase
	ExeCommand    string
	Timeout       time.Duration
}

// NewSandBoxRunner
func NewSandBoxRunner(path string, folder string, command string, file string, testcase []domain.Testcase, exeCommand string, submittedCode string, timeout time.Duration) *SandBoxRunner {
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

	execFile := utils.NewFileUtils(s.File)
	execFile.WriteCreateFile(s.SubmittedCode)

	return nil
}

func (s *SandBoxRunner) run() (*domain.CodeResult, error) {
	defer os.RemoveAll(s.Folder)

	stdout := ""

	before := time.Now()
	for _, t := range s.Testcase {
		output, err := secCom(s.ExeCommand, fmt.Sprintf("%s %s", s.Command, s.File), strings.NewReader(domain.CreateTestcase(t)), s.Timeout)

		if err != nil {
			return nil, err
		}

		bytesReader := bytes.NewReader(output)
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

		for i, s := range stdoutArray {
			if i < len(stdoutArray)-1 {
				stdout += s
			}
		}

		err = json.Unmarshal([]byte(stdoutArray[len(stdoutArray)-1]), &result)
		if result.Status == "fail" {
			after := time.Now()
			return &domain.CodeResult{
				Status:   "fail",
				Input:    t.Input,
				Expected: t.Output,
				Output:   stdout,
				Result:   result.Output,
				Time:     int(after.Unix() - before.Unix()),
			}, nil
		}
	}

	after := time.Now()
	if len(s.Testcase) == 1 {
		test := s.Testcase[0]
		return &domain.CodeResult{
			Status:   "success",
			Input:    test.Input,
			Expected: test.Output,
			Output:   stdout,
			Time:     int(after.Unix() - before.Unix()),
		}, nil

	}
	return &domain.CodeResult{
		Status: "success",
		Time:   int(after.Unix() - before.Unix()),
	}, nil
}

func secCom(execCommand, command string, input *strings.Reader, timeout time.Duration) ([]byte, error) {
	var policy = &seccomp.Policy{
		DefaultAction: seccomp.ActionAllow,
		Syscalls: []seccomp.SyscallGroup{
			{
				Action: seccomp.ActionErrno,
				Names: []string{
					"connect",
					"accept",
					"sendto",
					"recvfrom",
					"sendmsg",
					"recvmsg",
					"bind",
					"listen",
					"getpid",
					"kill",
					"fork",
				},
			},
		},
	}

	// Create a filter based on config.
	filter := seccomp.Filter{
		NoNewPrivs: noNewPrivs,
		Flag:       seccomp.FilterFlagTSync,
		Policy:     *policy,
	}

	// Load the BPF filter using the seccomp system call.
	if err := seccomp.LoadFilter(filter); err != nil {
		fmt.Fprintf(os.Stderr, "error loading filter: %v\n", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, execCommand, "-c", command)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	cmd.Stdin = input
	cmd.Env = []string{
		"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
	}
	if err := cmd.Run(); err != nil {
		return []byte{}, err
	}

	return out.Bytes(), nil
}
