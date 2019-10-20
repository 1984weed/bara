package remote

import (
	"bara/utils"
	"bytes"
	"fmt"
	"io"
	"log"
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
	_, err := os.Stat(s.SandboxFile)
	sandboxCommand := "" //s.SandboxFile
	if os.IsNotExist(err) {
		sandboxCommand = ""
	}

	log.Println(fmt.Sprintf("%s %s %s", sandboxCommand, s.Command, s.File))

	// return exec.Command("cat", s.TestcaseFile).Output()
	c1 := exec.Command("cat", s.TestcaseFile)
	c2 := exec.Command(s.ExeCommand, "-c", fmt.Sprintf("%s %s %s", sandboxCommand, s.Command, s.File))
	log.Println(c2)
	c2.Env = []string{}

	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	var b2 bytes.Buffer
	c2.Stdout = &b2

	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()

	defer os.RemoveAll(s.Folder)

	log.Println("========================")
	log.Println("out", b2.String())
	log.Println("========================")
	return b2.Bytes(), nil
}
