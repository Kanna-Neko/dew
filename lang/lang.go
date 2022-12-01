package lang

import (
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

type Lang struct {
	IsComplieLang bool

	// compile file name must is cat
	CompileCode func(codefile string) *exec.Cmd
	RunCode     func(codefile string) *exec.Cmd

	ProgramTypeId string
	Codefile      string
	Name          string
}

var LangSlice = []string{"c", "c++", "c++17", "go", "python3"}
var OriginLangDic map[string]ExportLang = map[string]ExportLang{
	"c": {
		IsComplieLang:  true,
		CompileCommand: "gcc $codefile -o cat",
		RunCommand:     "./cat",
		ProgramTypeId:  "43",
		Codefile:       "dew.c",
		Name:           "GNU GCC C11 5.1.0",
	},
	"c++": {
		IsComplieLang:  true,
		CompileCommand: "g++ $codefile -o cat",
		RunCommand:     "./cat",
		ProgramTypeId:  "61",
		Codefile:       "dew.cpp",
		Name:           "GNU G++17 9.2.0 (64 bit, msys 2)",
	},
	"c++17": {
		IsComplieLang:  true,
		CompileCommand: "g++ $codefile -o cat -std=c++17",
		RunCommand:     "./cat",
		ProgramTypeId:  "54",
		Codefile:       "dew.cpp",
		Name:           "GNU G++17 7.3.0",
	},
	"go": {
		IsComplieLang:  true,
		CompileCommand: "go build -o cat $codefile",
		RunCommand:     "./cat",
		ProgramTypeId:  "32",
		Codefile:       "dew.go",
		Name:           "Go 1.19",
	},
	"python3": {
		IsComplieLang:  false,
		CompileCommand: " ",
		RunCommand:     "python3 $codefile",
		ProgramTypeId:  "31",
		Codefile:       "dew.py",
		Name:           "Python 3.8.10",
	},
}

//language dictionary
var LangDic map[string]Lang = map[string]Lang{
	// "c": {
	// 	IsComplieLang: true,
	// 	CompileCode:   func(codefile string) *exec.Cmd { return exec.Command("gcc", codefile, "-o", "cat") },
	// 	RunCode:       func(codefile string) *exec.Cmd { return exec.Command("./cat") },
	// 	ProgramTypeId: "43",
	// 	Codefile:      "dew.c",
	// 	Name:          "GNU GCC C11 5.1.0",
	// },
	// "c++": {
	// 	IsComplieLang: true,
	// 	CompileCode:   func(codefile string) *exec.Cmd { return exec.Command("g++", codefile, "-o", "cat") },
	// 	RunCode:       func(codefile string) *exec.Cmd { return exec.Command("./cat") },
	// 	ProgramTypeId: "61",
	// 	Codefile:      "dew.cpp",
	// 	Name:          "GNU G++17 9.2.0 (64 bit, msys 2)",
	// },
	// "c++17": {
	// 	IsComplieLang: true,
	// 	CompileCode:   func(codefile string) *exec.Cmd { return exec.Command("g++", codefile, "-o", "cat", "-std=c++17") },
	// 	RunCode:       func(codefile string) *exec.Cmd { return exec.Command("./cat") },
	// 	ProgramTypeId: "54",
	// 	Codefile:      "dew.cpp",
	// 	Name:          "GNU G++17 7.3.0",
	// },
	// "go": {
	// 	IsComplieLang: true,
	// 	CompileCode:   func(codefile string) *exec.Cmd { return exec.Command("go", "build", "-o", "cat", codefile) },
	// 	RunCode:       func(codefile string) *exec.Cmd { return exec.Command("./cat") },
	// 	ProgramTypeId: "32",
	// 	Codefile:      "dew.go",
	// 	Name:          "Go 1.19",
	// },
	// "python3": {
	// 	IsComplieLang: false,
	// 	CompileCode:   func(codefile string) *exec.Cmd { return exec.Command("") },
	// 	RunCode:       func(codefile string) *exec.Cmd { return exec.Command("python3", codefile) },
	// 	ProgramTypeId: "31",
	// 	Codefile:      "dew.py",
	// 	Name:          "Python 3.8.10",
	// },
}

type ExportLang struct {
	IsComplieLang  bool   `json:"isCompileLang"`
	CompileCommand string `json:"compileCommand"`
	RunCommand     string `json:"runCommand"`
	ProgramTypeId  string `json:"programTypeId"`
	Codefile       string `json:"codefile"`
	Name           string `json:"name"`
}

func ImportLangDic() {
	viper.SetConfigFile("./codeforces/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("please use init command first, error : %s\n", err)
	}
	var info = viper.GetStringMap("language")
	for k := range info {
		var data ExportLang
		data.Codefile = viper.GetString("language." + k + ".codefile")
		data.IsComplieLang = viper.GetBool("language." + k + ".isCompileLang")
		data.CompileCommand = viper.GetString("language." + k + ".compileCommand")
		data.RunCommand = viper.GetString("language." + k + ".runCommand")
		data.ProgramTypeId = viper.GetString("language." + k + ".programTypeId")
		data.Name = viper.GetString("language." + k + ".name")
		data.CompileCommand = strings.Trim(data.CompileCommand, " ")
		data.RunCommand = strings.Trim(data.RunCommand, " ")
		var compileSlice = strings.Split(data.CompileCommand, " ")
		var runSlice = strings.Split(data.RunCommand, " ")
		var compile = func(codefile string) *exec.Cmd {
			for k, v := range compileSlice {
				if v == "$codefile" {
					compileSlice[k] = codefile
				}
			}
			return exec.Command(compileSlice[0], compileSlice[1:]...)
		}
		var run = func(codefile string) *exec.Cmd {
			if data.RunCommand == "" {
				return exec.Command("")
			}
			for k, v := range runSlice {
				if v == "$codefile" {
					runSlice[k] = codefile
				}
			}
			return exec.Command(runSlice[0], runSlice[1:]...)
		}
		LangDic[k] = Lang{
			IsComplieLang: data.IsComplieLang,
			CompileCode:   compile,
			RunCode:       run,
			ProgramTypeId: data.ProgramTypeId,
			Codefile:      data.Codefile,
			Name:          data.Name,
		}
	}
}
