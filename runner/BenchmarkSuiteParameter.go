package runner

import (
	"fmt"

	bench "github.com/pip-benchmark/pip-benchmark-go/benchmark"
)

type BenchmarkSuiteParameter struct {
	*bench.Parameter
	originalParameter *bench.Parameter
}

func NewBenchmarkSuiteParameter(suite *BenchmarkSuiteInstance, originalParameter *bench.Parameter) *BenchmarkSuiteParameter {
	c := BenchmarkSuiteParameter{}

	c.Parameter = bench.NewParameter(fmt.Sprintf("%s.%s", suite.Name(), originalParameter.Name()),
		originalParameter.Description(), originalParameter.DefaultValue(), "BenchmarkSuiteParameter")
	c.originalParameter = originalParameter
	c.IParameter = &c
	return &c
}

func (c *BenchmarkSuiteParameter) GetValue() string {
	return c.originalParameter.Value()
}

func (c *BenchmarkSuiteParameter) SetValue(value string) {
	c.originalParameter.SetValue(value)
}
