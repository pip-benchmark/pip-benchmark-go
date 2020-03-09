package parameters

import (
	"fmt"

	bench "github.com/pip-benchmark/pip-benchmark-go/benchmark"
	benchmarks "github.com/pip-benchmark/pip-benchmark-go/runner/benchmarks"
	benchconv "github.com/pip-benchmark/pip-benchmark-go/utilities"
)

type BenchmarkSelectedParameter struct {
	*bench.Parameter
	benchmark *benchmarks.BenchmarkInstance
}

func NewBenchmarkSelectedParameter(benchmark *benchmarks.BenchmarkInstance) *BenchmarkSelectedParameter {
	c := BenchmarkSelectedParameter{}
	c.Parameter = bench.NewParameter(fmt.Sprintf("%s.%s.Selected", benchmark.Suite().Name(), benchmark.Name()),
		fmt.Sprintf("Selecting benchmark %s in suite %s", benchmark.Name(), benchmark.Suite().Name()),
		"true")
	c.benchmark = benchmark
	return &c
}

func (c *BenchmarkSelectedParameter) GetValue() string {
	return benchconv.Converter.BooleanToString(c.benchmark.IsSelected())
}

func (c *BenchmarkSelectedParameter) SetValue(value string) {
	c.benchmark.Select(benchconv.Converter.StringToBoolean(value, false))
}
