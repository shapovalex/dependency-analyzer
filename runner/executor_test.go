package runner

import (
	"github.com/shapovalex/depAnalyzer/processor"
	"github.com/shapovalex/depAnalyzer/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCallPypiComparatorOnSingleProcessor(t *testing.T) {
	pypiComparator := &test.MockProcessor{Manager: "pypi", Operation: "compare"}
	processors := []processor.Processor{pypiComparator}
	executor := NewExecutor(processors)
	params := processor.Params{
		DependencyManager: "pypi",
		InputFiles:        "../test/test_input.txt",
		Operation:         "compare",
	}

	executor.Execute(params)
	assert.True(t, pypiComparator.Called)
	assert.Equal(t, "pypi", pypiComparator.Params.DependencyManager)
	assert.Equal(t, "compare", pypiComparator.Params.Operation)
	assert.Equal(t, "../test/test_input.txt", pypiComparator.Params.InputFiles)

}

func TestShouldCallPypiComparatorOnMultipleProcessors(t *testing.T) {

	pypiComparator := &test.MockProcessor{Manager: "pypi", Operation: "compare"}
	mavenComparator := &test.MockProcessor{Manager: "maven", Operation: "compare"}
	processors := []processor.Processor{mavenComparator, pypiComparator}
	executor := NewExecutor(processors)
	params := processor.Params{
		DependencyManager: "pypi",
		InputFiles:        "../test/test_input.txt",
		Operation:         "compare",
	}

	executor.Execute(params)
	assert.Equal(t, "../test/test_input.txt", pypiComparator.Params.InputFiles)
	assert.False(t, mavenComparator.Called)
}

func TestShouldCallPypiComparatorBasedOnOperation(t *testing.T) {

	pypiComparator1 := &test.MockProcessor{Manager: "pypi", Operation: "compare"}
	pypiComparator2 := &test.MockProcessor{Manager: "pypi", Operation: "licenses"}
	processors := []processor.Processor{pypiComparator1, pypiComparator2}
	executor := NewExecutor(processors)
	params := processor.Params{
		DependencyManager: "pypi",
		InputFiles:        "",
		Operation:         "licenses",
	}

	executor.Execute(params)
	assert.True(t, pypiComparator2.Called)
	assert.False(t, pypiComparator1.Called)
}

func TestShouldPassFileContentInParam(t *testing.T) {
	pypiComparator := &test.MockProcessor{Manager: "pypi", Operation: "compare"}
	processors := []processor.Processor{pypiComparator}
	executor := NewExecutor(processors)
	params := processor.Params{
		DependencyManager: "pypi",
		InputFiles:        "../test/test_input.txt,../test/test_input2.txt",
		Operation:         "compare",
	}

	executor.Execute(params)
	assert.Equal(t, "Line1\nLine2", pypiComparator.Params.InputFilesContent[0])
	assert.Equal(t, "Line3\nLine4", pypiComparator.Params.InputFilesContent[1])
}
