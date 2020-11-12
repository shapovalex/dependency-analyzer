package runner

import (
	"depAnalyzer/src/processor"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockedComparator struct {
	mock.Mock
}

func (m *MockedComparator) Process(params processor.Params) {
	m.Called(params)
}

func (m *MockedComparator) GetSupportedDependencyManager() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockedComparator) GetSupportedOperation() string {
	args := m.Called()
	return args.String(0)
}

func TestShouldCallPypiComparatorOnSingleProcessor(t *testing.T) {

	pypiComparator := new(MockedComparator)
	processors := []processor.Processor{pypiComparator}
	executor := NewExecutor(processors)
	params := processor.Params{
		DependencyManager: "pypi",
		InputFiles:        "requirements1.txt,requirements2.txt",
		Operation:         "compare",
	}
	pypiComparator.On("Process", mock.Anything).Return()
	pypiComparator.On("GetSupportedDependencyManager").Return("pypi")
	pypiComparator.On("GetSupportedOperation").Return("compare")

	executor.Execute(params)
	pypiComparator.AssertCalled(t, "Process", params)
}

func TestShouldCallPypiComparatorOnMultipleProcessors(t *testing.T) {

	pypiComparator := new(MockedComparator)
	mavenComparator := new(MockedComparator)
	processors := []processor.Processor{mavenComparator, pypiComparator}
	executor := NewExecutor(processors)
	params := processor.Params{
		DependencyManager: "pypi",
		InputFiles:        "requirements1.txt,requirements2.txt",
		Operation:         "compare",
	}
	pypiComparator.On("Process", mock.Anything).Return()
	pypiComparator.On("GetSupportedDependencyManager").Return("pypi")
	pypiComparator.On("GetSupportedOperation").Return("compare")

	mavenComparator.On("Process", mock.Anything).Return()
	mavenComparator.On("GetSupportedDependencyManager").Return("maven")
	mavenComparator.On("GetSupportedOperation").Return("compare")

	executor.Execute(params)
	pypiComparator.AssertCalled(t, "Process", params)
	mavenComparator.AssertNotCalled(t, "Process", params)
}

func TestShouldCallPypiComparatorBasedOnOperation(t *testing.T) {

	pypiComparator1 := new(MockedComparator)
	pypiComparator2 := new(MockedComparator)
	processors := []processor.Processor{pypiComparator1, pypiComparator2}
	executor := NewExecutor(processors)
	params := processor.Params{
		DependencyManager: "pypi",
		InputFiles:        "requirements1.txt,requirements2.txt",
		Operation:         "licenses",
	}
	pypiComparator1.On("Process", mock.Anything).Return()
	pypiComparator1.On("GetSupportedDependencyManager").Return("pypi")
	pypiComparator1.On("GetSupportedOperation").Return("compare")

	pypiComparator2.On("Process", mock.Anything).Return()
	pypiComparator2.On("GetSupportedDependencyManager").Return("pypi")
	pypiComparator2.On("GetSupportedOperation").Return("licenses")

	executor.Execute(params)
	pypiComparator1.AssertNotCalled(t, "Process", params)
	pypiComparator2.AssertCalled(t, "Process", params)
}
