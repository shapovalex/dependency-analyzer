package main

import (
	"depAnalyzer/src/processor"
	"depAnalyzer/src/processor/comparator"
	"depAnalyzer/src/processor/license"
	"depAnalyzer/src/runner"
	"flag"
)

func main() {
	var dFlag = flag.String("d", "", "Specify dependency manager to use. Allowed values: pypi")
	var fFlag = flag.String("f", "", "Specify file/files to analyze")
	var oFlag = flag.String("o", "", "Specify operation to perform")
	var rFlag = flag.String("r", "result.txt", "Specify file to write results")
	flag.Parse()

	params := processor.Params{}
	params.Operation = *oFlag
	params.InputFiles = *fFlag
	params.DependencyManager = *dFlag
	params.OutputFiles = *rFlag

	pypiComparator := new(comparator.PyPiComparator)
	pypiLicense := new(license.PyPiLicense)

	executor := runner.NewExecutor([]processor.Processor{
		pypiComparator,
		pypiLicense,
	})

	executor.Execute(params)
}
