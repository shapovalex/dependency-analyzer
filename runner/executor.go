package runner

import (
	"depAnalyzer/processor"
	"fmt"
)

type Executor struct {
	processors []processor.Processor
}

func (e Executor) Execute(params processor.Params) {
	processed := false
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
