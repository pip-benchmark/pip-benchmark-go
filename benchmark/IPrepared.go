package benchmark

type IPrepared interface {
	SetUp() error
	TearDown() error
}
