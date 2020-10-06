package runner

type IExecutionStrategy interface {
	Start() error
	StopExecution() error
	IsStopped() bool
}
