package benchmark

type IExecutable interface {
	IPrepared
	Execute() error
}
