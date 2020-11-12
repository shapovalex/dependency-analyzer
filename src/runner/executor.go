package runner

import "depAnalyzer/src/processor"

type Executor struct {
	processors []processor.Processor
}

func (e Executor) execute(params processor.Params) {
	for _, processorItem := range e.processors {
		if params.DependencyManager == processorItem.GetSupportedDependencyManager() {
			processorItem.Process(params)
		}
	}
}

func NewExecutor(processors []processor.Processor) Executor {
	return Executor{processors: processors}
}
