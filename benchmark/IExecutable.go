package benchmark

type IExecutable interface {
	Execute() error
	SetUp() error
	TearDown() error
}
