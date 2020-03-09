package execution

type IExecutionStrategy interface {
	Start() error
	Stop() error
	IsStopped() bool
}
