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

var LangSlice = []string{"c", "c++", "c++17", "go"}

//language dictionary
var LangDic map[string]Lang = map[string]Lang{
	"c": {
		IsComplieLang:    true,
		CompileCode:      func(codefile string) *exec.Cmd { return exec.Command("gcc", codefile, "-o", "cat") },
		RunCode:          func(codefile string) *exec.Cmd { return exec.Command("./cat") },
		ProgramTypeId:    "43",
		OriginalCodefile: "dew.c",
		Name:             "GNU GCC C11 5.1.0",
	},
	"c++": {
		IsComplieLang:    true,
		CompileCode:      func(codefile string) *exec.Cmd { return exec.Command("g++", codefile, "-o", "cat") },
		RunCode:          func(codefile string) *exec.Cmd { return exec.Command("./cat") },
		ProgramTypeId:    "61",
		OriginalCodefile: "dew.cpp",
		Name:             "GNU G++17 9.2.0 (64 bit, msys 2)",
	},
	"c++17": {
		IsComplieLang:    true,
		CompileCode:      func(codefile string) *exec.Cmd { return exec.Command("g++", codefile, "-o", "cat", "-std=c++17") },
		RunCode:          func(codefile string) *exec.Cmd { return exec.Command("./cat") },
		ProgramTypeId:    "54",
		OriginalCodefile: "dew.cpp",
		Name:             "GNU G++17 7.3.0",
	},
	"go": {
		IsComplieLang:    true,
		CompileCode:      func(codefile string) *exec.Cmd { return exec.Command("go", "build", "-o", "cat", codefile) },
		RunCode:          func(codefile string) *exec.Cmd { return exec.Command("./cat") },
		ProgramTypeId:    "32",
		OriginalCodefile: "dew.go",
		Name:             "Go 1.19",
	},
}
