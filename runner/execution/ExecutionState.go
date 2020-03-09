package execution

type ExecutionState int

const (
	Initial ExecutionState = iota
	Running
	Completed
)
