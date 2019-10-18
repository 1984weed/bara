package remote

type CompileLanguage string

const (
	JavaScript CompileLanguage = "javascript"
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

var CodeCompileMap = map[CompileLanguage]CompileInfo{
	JavaScript: {
		FileName:    "file.js",
		Command:     "node",
		PrepareCode: nodeJsTemplate,
	},
}

var MachineExecMap = map[MachineType]string{
	Container: "ash",
	LocalMac:  "bash",
}
