package parameters

import (
	"fmt"

	bench "github.com/pip-benchmark/pip-benchmark-go/benchmark"
	benchmarks "github.com/pip-benchmark/pip-benchmark-go/runner/benchmarks"
)

type BenchmarkSuiteParameter struct {
	*bench.Parameter
	originalParameter *bench.Parameter
}

func NewBenchmarkSuiteParameter(suite *benchmarks.BenchmarkSuiteInstance, originalParameter *bench.Parameter) *BenchmarkSuiteParameter {
	c := BenchmarkSuiteParameter{}

	c.Parameter = bench.NewParameter(fmt.Sprintf("%s.%s", suite.Name(), originalParameter.Name()),
		originalParameter.Description(), originalParameter.DefaultValue())
	c.originalParameter = originalParameter

	return &c
}

func (c *BenchmarkSuiteParameter) GetValue() string {
	return c.originalParameter.Value()
}

func (c *BenchmarkSuiteParameter) SetValue(value string) {
	c.originalParameter.SetValue(value)
}
