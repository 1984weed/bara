package executor

import (
	"bara/model"
)

type MachineType string

const (
	Container MachineType = "container"
	LocalMac  MachineType = "localMac"
)

type CompileInfo struct {
	FileName    string
	Command     string
	PrepareCode string
}

var CodeCompileMap = map[model.CodeLanguageSlug]CompileInfo{
	model.JavaScript: {
		FileName:    "file.js",
		Command:     "node",
		PrepareCode: Node,
	},
}

var MachineExecMap = map[MachineType]string{
	Container: "sh",
	LocalMac:  "bash",
}
