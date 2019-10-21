package remote

import (
	"bara/utils"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
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
	Testcase      string
	ExeCommand    string
	Timeout       int
	Language      CompileLanguage
}

func NewSandBoxRunner(path, folder, command, file, testcase, exeCommand, submittedCode string, timeout int, language CompileLanguage) *SandBoxRunner {
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
		Language:      language,
	}
}

func (s *SandBoxRunner) Exec() ([]byte, error) {
	err := s.prepare()
	if err != nil {
		return []byte{}, err
	}
	return s.run()
}

func (s *SandBoxRunner) prepare() error {
	err := os.MkdirAll(s.Folder, os.ModePerm)

	if err != nil {
		return err
	}
	testcaseFile := utils.NewFileUtils(s.TestcaseFile)
	err = testcaseFile.WriteCreateFile(s.Testcase)

	if err != nil {
		return err
	}

	execFile := utils.NewFileUtils(s.File)
	execFile.WriteCreateFile(s.SubmittedCode)

	return nil
}

func (s *SandBoxRunner) run() ([]byte, error) {
	sandboxCommand := s.SandboxFile
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)

	if _, err := os.Stat(s.SandboxFile); os.IsNotExist(err) {
		sandboxCommand = ""
		return exec.CommandContext(ctx, s.ExeCommand, "-c", fmt.Sprintf("cat %s | %s %s %s", s.TestcaseFile, sandboxCommand, s.Command, s.File)).Output()
	}

	u, _ := user.Lookup("execUser")

	uid, _ := strconv.Atoi(u.Uid)
	gid, _ := strconv.Atoi(u.Gid)

	log.Println(fmt.Sprintf("%s %s %s", sandboxCommand, s.Command, s.File))
	defer cancel()
	defer os.RemoveAll(s.Folder)

	cmd := exec.CommandContext(ctx, s.ExeCommand, "-c", fmt.Sprintf("cat %s | %s %s %s", s.TestcaseFile, sandboxCommand, s.Command, s.File))
	cmd.Env = []string{
		"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
	}
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid),
		Gid: uint32(gid)}

	return cmd.Output()
}
