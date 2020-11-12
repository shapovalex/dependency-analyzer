package comparator

import "depAnalyzer/src/processor"

type PyPiComparator struct {
}

func (p PyPiComparator) GetSupportedDependencyManager() string {
	return "pypi"
}

func (p PyPiComparator) Process(params processor.Params) {

}
