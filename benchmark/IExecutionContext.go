package benchmark

type IExecutionContext interface {
	//parameters interface{}
	IncrementCounter(increment int64)
	SendMessage(message string)
	ReportError(err error)
	//isStopped bool
	Stop()
}
