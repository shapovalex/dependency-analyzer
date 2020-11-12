package comparator

import (
	"depAnalyzer/src/helper"
	"depAnalyzer/src/processor"
)

type PyPiComparator struct {
}

func (p PyPiComparator) GetSupportedDependencyManager() string {
	return "pypi"
}

func (p PyPiComparator) Process(params processor.Params) {

}

func (p PyPiComparator) GetSupportedOperation() string {
	return "compare"
}

func (p PyPiComparator) compare(list1 []string, list2 []string) []string {
	return helper.DeleteIntersectionElements(list2, list1)
}
