package runner

type IExecutionStrategy interface {
	Start() error
	Stop() error
	IsStopped() bool
}
