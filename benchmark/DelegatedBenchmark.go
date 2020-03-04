package benchmark

type DelegatedBenchmark struct {
	*Benchmark
	executeCallback func() error
}

func NewDelegatedBenchmark(name string, description string,
	executeCallback func() error) *DelegatedBenchmark {

	if executeCallback == nil {
		panic("DelegatedBenchmark: ExecuteCallback can't be nil")
	}
	c := DelegatedBenchmark{
		Benchmark:       NewBenchmark(name, description),
		executeCallback: executeCallback,
	}
	c.Benchmark.IExecutable = &c
	return &c
}

func (c *DelegatedBenchmark) Execute() error {
	return c.executeCallback()
}
