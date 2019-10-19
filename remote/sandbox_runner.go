package remote

import (
	"bara/utils"
	"fmt"
	"os"
	"os/exec"
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
	var _, err = os.Stat(s.SandboxFile)
	sandboxCommand := s.SandboxFile
	if os.IsNotExist(err) {
		sandboxCommand = ""
	}
	inputCommand := fmt.Sprintf(`cat %s | %s %s %s`, s.TestcaseFile, sandboxCommand, s.Command, s.File)

	defer os.RemoveAll(s.Folder)
	return exec.Command(s.ExeCommand, "-c", inputCommand).Output()
}
