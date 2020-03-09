package benchmark

type IExecutionContext interface {
	//parameters interface{}
	IncrementCounter(increment int)
	SendMessage(message string)
	ReportError(err error)
	//isStopped bool
	Stop()
}
