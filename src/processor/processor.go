package processor

type Processor interface {
	Process(params Params)
	GetSupportedDependencyManager() string
	GetSupportedOperation() string
}
