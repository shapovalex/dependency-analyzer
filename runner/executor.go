package runner

import (
	"fmt"
	"github.com/shapovalex/depAnalyzer/processor"
	"io/ioutil"
	"os"
	"strings"
)

type Executor struct {
	processors []processor.Processor
}

func readFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Unable to read file", filename)
		os.Exit(1)
	}
	return string(content)
}

func readContent(files string) []string {
	inputFiles := strings.Split(files, ",")
	var result []string
	for _, fileName := range inputFiles {
		if fileName != "" {
			content := readFile(fileName)
			result = append(result, content)
		}
	}
	return result
}

func (e Executor) Execute(params processor.Params) {
	processed := false
	params.InputFilesContent = readContent(params.InputFiles)
	for _, processorItem := range e.processors {
		if params.DependencyManager == processorItem.GetSupportedDependencyManager() &&
			params.Operation == processorItem.GetSupportedOperation() {
			processed = true
			processorItem.Process(params)
			fmt.Println("Request successfully processed")
		}
	}
	if !processed {
		fmt.Println("Unable to find operation")
	}
}

func NewExecutor(processors []processor.Processor) Executor {
	return Executor{processors: processors}
}
