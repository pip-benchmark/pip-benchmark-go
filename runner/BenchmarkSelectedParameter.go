package runner

import (
	"fmt"

	bench "github.com/pip-benchmark/pip-benchmark-go/benchmark"
	util "github.com/pip-benchmark/pip-benchmark-go/utilities"
)

type BenchmarkSelectedParameter struct {
	*bench.Parameter
	benchmark *BenchmarkInstance
}

func NewBenchmarkSelectedParameter(benchmark *BenchmarkInstance) *BenchmarkSelectedParameter {
	c := BenchmarkSelectedParameter{}
	c.Parameter = bench.NewParameter(fmt.Sprintf("%s.%s.Selected", benchmark.Suite().Name(), benchmark.Name()),
		fmt.Sprintf("Selecting benchmark %s in suite %s", benchmark.Name(), benchmark.Suite().Name()),
		"true", "BenchmarkSelectedParameter")
	c.benchmark = benchmark
	return &c
}

func (c *BenchmarkSelectedParameter) GetValue() string {
	return util.Converter.BooleanToString(c.benchmark.IsSelected())
}

func (c *BenchmarkSelectedParameter) SetValue(value string) {
	c.benchmark.Select(util.Converter.StringToBoolean(value, false))
}
