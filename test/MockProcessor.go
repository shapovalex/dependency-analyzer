package test

import (
	"github.com/shapovalex/depAnalyzer/processor"
)

type MockProcessor struct {
	Manager   string
	Operation string
	Params    processor.Params
	Called    bool
}

func (m *MockProcessor) Process(params processor.Params) {
	m.Params = params
	m.Called = true
}

func (m MockProcessor) GetSupportedDependencyManager() string {
	return m.Manager
}

func (m MockProcessor) GetSupportedOperation() string {
	return m.Operation
}
