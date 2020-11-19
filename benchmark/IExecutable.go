package benchmark

type IExecutable interface {
	SetUp() error
	TearDown() error
	Execute() error
}
