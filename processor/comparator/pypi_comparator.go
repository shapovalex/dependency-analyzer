package comparator

import (
	"fmt"
	"github.com/shapovalex/depAnalyzer/helper"
	"github.com/shapovalex/depAnalyzer/processor"
	"os"
	"strings"
)

type PyPiComparator struct {
}

func (p PyPiComparator) GetSupportedDependencyManager() string {
	return "pip"
}

func (p PyPiComparator) Process(params processor.Params) {
	inputFiles := strings.Split(params.InputFiles, ",")
	lines1, errR1 := helper.ReadLines(inputFiles[0])
	lines2, errR2 := helper.ReadLines(inputFiles[1])
	if errR1 != nil || errR2 != nil {
		fmt.Println("Unable to read input files")
		os.Exit(1)
	}
	result := p.compare(lines1, lines2)
	errW := helper.WriteLines(result, params.OutputFiles)
	if errW != nil {
		fmt.Println("Unable to write to output files")
		os.Exit(1)
	}
}

func (p PyPiComparator) GetSupportedOperation() string {
	return "compare"
}

func (p PyPiComparator) compare(list1 []string, list2 []string) []string {
	return helper.DeleteIntersectionElements(list2, list1)
}
