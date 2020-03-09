package parameters

import (
	"fmt"

	bench "github.com/pip-benchmark/pip-benchmark-go/benchmark"
	benchmarks "github.com/pip-benchmark/pip-benchmark-go/runner/benchmarks"
	benchconv "github.com/pip-benchmark/pip-benchmark-go/utilities"
)

type BenchmarkProportionParameter struct {
	*bench.Parameter
	benchmark *benchmarks.BenchmarkInstance
}

func NewBenchmarkProportionParameter(benchmark *benchmarks.BenchmarkInstance) *BenchmarkProportionParameter {

	c := BenchmarkProportionParameter{}
	c.Parameter = bench.NewParameter(fmt.Sprintf("%s.%s.Proportion", benchmark.Suite().Name(), benchmark.Name()),
		fmt.Sprintf("Sets execution proportion for benchmark %s in suite %s", benchmark.Name(), benchmark.Suite().Name()),
		"100")
	c.benchmark = benchmark

	return &c
}

func (c *BenchmarkProportionParameter) GetValue() string {
	return benchconv.Converter.IntegerToString(c.benchmark.GetProportion())
}

func (c *BenchmarkProportionParameter) SetValue(value string) {
	c.benchmark.SetProportion(benchconv.Converter.StringToInteger(value, 100))
}
