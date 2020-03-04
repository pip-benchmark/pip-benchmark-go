package benchmark

import (
	cerr "github.com/pip-services3-go/pip-services3-commons-go/errors"
)

type PassiveBenchmark struct {
	*Benchmark
}

func NewPassiveBenchmark(name string, description string) *PassiveBenchmark {
	c := PassiveBenchmark{
		Benchmark: NewBenchmark(name, description),
	}
	c.Benchmark.IExecutable = &c
	return &c
}

func (c *PassiveBenchmark) Execute() error {
	return cerr.NewError("Active measurement via Execute is not allow for passive benchmarks")
}
