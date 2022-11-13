package testmanager

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const (
	managerFilePath = "./codeforces/testManager.json"
	testFilesDir    = "./codeforces/testFiles/"
	capacity        = 50
)

func init() {
	err := os.MkdirAll(testFilesDir, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func StoreProblemTest(problem string, tests Testcases) {
	data, err := json.Marshal(tests)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(testFilesDir+problem+".json", data, 0777)
	if err != nil {
		log.Fatal(err)
	}
}
func ExtractTestcase(problem string) Testcases {
	var tests Testcases
	data, err := ioutil.ReadFile(testFilesDir + problem + ".json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &tests)
	if err != nil {
		log.Fatal(err)
	}
	return tests
}

func ManagerPush(manager Manager, problem string) Manager {
	for i, v := range manager.Problems {
		if v == problem {
			manager.Problems = append(manager.Problems[:i], manager.Problems[i+1:]...)
			manager.Problems = append([]string{problem}, manager.Problems...)
			return manager
		}
	}
	manager.Problems = append([]string{problem}, manager.Problems...)
	return manager
}

func ManagerDeleteExtra(manager Manager) Manager {
	for i := 50; i < len(manager.Problems); i++ {
		os.Remove(testFilesDir + manager.Problems[i] + ".json")
	}
	if len(manager.Problems) > 50 {
		manager.Problems = manager.Problems[0:50]
	}
	return manager
}
func IsTestcaseExist(problem string) bool {
	_, err := os.Stat(testFilesDir + problem + ".json")
	return !os.IsNotExist(err)
}

func ExtractManager() Manager {
	var res Manager
	res.Problems = make([]string, 0)
	_, err := os.Stat(managerFilePath)
	if os.IsNotExist(err) {
		return res
	}
	data, err := ioutil.ReadFile(managerFilePath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func StoreManager(data Manager) {
	js, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(managerFilePath, js, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

type Testcase struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type Testcases struct {
	Tests []Testcase `json:"tests"`
}

type Manager struct {
	Problems []string `json:"problems"`
}
