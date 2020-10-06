package benchmark

type IExecutionContext interface {
	GetParameters() map[string]*Parameter
	IncrementCounter(increment int)
	SendMessage(message string)
	ReportError(err error)
	//isStopped bool
	Stop()
}
