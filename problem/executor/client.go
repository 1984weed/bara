package executor

import (
	"bara/model"
	"bara/problem/domain"
	"bara/utils"
	"fmt"
	"log"
	"os"
	"time"
)

type Client interface {
	Exec(codeLanguage model.CodeLanguageSlug, typedCode string, testcase []string, functionName string) (*domain.CodeResult, error)
}

type executor struct {
	withoutContainer bool
	timeoutSecond    time.Duration
}

func NewExecutorClient(withoutContainer bool, timeoutSecond time.Duration) Client {
	return &executor{
		withoutContainer,
		timeoutSecond,
	}
}

func (e *executor) Exec(codeLanguage model.CodeLanguageSlug, typedCode string, testcase []string, functionName string) (*domain.CodeResult, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, nil
	}

	codeCompile := CodeCompileMap[codeLanguage]
	var machine MachineType
	if e.withoutContainer {
		machine = LocalMac
	} else {
		machine = Container
	}
	machineType := MachineExecMap[machine]
	execStr := fmt.Sprintf(codeCompile.PrepareCode, typedCode, functionName)
	log.Println(machineType)

	sandbox := NewSandBoxRunner(dir, fmt.Sprintf(`folder-%s`, utils.RandomString(10)), codeCompile.Command, codeCompile.FileName, testcase, machineType, execStr, e.timeoutSecond)
	result, err := sandbox.Exec()

	if err != nil {
		return nil, err
	}

	// log.Print("output from command", string(out))

	// bytesReader := bytes.NewReader(out)
	// reader := bufio.NewReader(bytesReader)
	// var result domain.CodeResult
	// stdoutArray := []string{}

	// for {
	// 	line, err := reader.ReadString('\n')

	// 	if err == io.EOF {
	// 		break
	// 	}

	// 	stdoutArray = append(stdoutArray, fmt.Sprintf("%s", line))

	// }

	// stdout := ""

	// for i, s := range stdoutArray {
	// 	if i < len(stdoutArray)-1 {
	// 		stdout += s
	// 	}
	// }

	// err = json.Unmarshal([]byte(stdoutArray[len(stdoutArray)-1]), &result)

	// result.Output = stdout

	return result, nil
}
