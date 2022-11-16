package lang

import "os/exec"

type Lang struct {
	IsComplieLang bool

	// compile file name must is cat
	CompileCode func(codefile string) *exec.Cmd
	RunCode     func(codefile string) *exec.Cmd

	ProgramTypeId    string
	OriginalCodefile string
	Name             string
}

var LangSlice = []string{"c++", "c++17"}

//language dictionary
var LangDic map[string]Lang = map[string]Lang{
	"c++": {
		IsComplieLang:    true,
		CompileCode:      func(codefile string) *exec.Cmd { return exec.Command("g++", codefile, "-o", "cat") },
		RunCode:          func(codefile string) *exec.Cmd { return exec.Command("./cat") },
		ProgramTypeId:    "61",
		OriginalCodefile: "dew.cpp",
		Name:             "GNU C++17 (64)",
	},
	"c++17": {
		IsComplieLang:    true,
		CompileCode:      func(codefile string) *exec.Cmd { return exec.Command("g++", codefile, "-o", "cat", "-std=c++17") },
		RunCode:          func(codefile string) *exec.Cmd { return exec.Command("./cat") },
		ProgramTypeId:    "54",
		OriginalCodefile: "dew.cpp",
		Name:             "GNU G++17 7.3.0",
	},
}
