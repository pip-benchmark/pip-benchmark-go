package runner

import (
	"fmt"

	bench "github.com/pip-benchmark/pip-benchmark-go/benchmark"
	util "github.com/pip-benchmark/pip-benchmark-go/utilities"
)

type BenchmarkProportionParameter struct {
	*bench.Parameter
	benchmark *BenchmarkInstance
}

func NewBenchmarkProportionParameter(benchmark *BenchmarkInstance) *BenchmarkProportionParameter {

	c := BenchmarkProportionParameter{}
	c.Parameter = bench.NewParameter(fmt.Sprintf("%s.%s.Proportion", benchmark.Suite().Name(), benchmark.Name()),
		fmt.Sprintf("Sets execution proportion for benchmark %s in suite %s", benchmark.Name(), benchmark.Suite().Name()),
		"100", "BenchmarkProportionParameter")
	c.benchmark = benchmark

	return &c
}

func (c *BenchmarkProportionParameter) GetValue() string {
	return util.Converter.IntegerToString(c.benchmark.GetProportion())
}

func (c *BenchmarkProportionParameter) SetValue(value string) {
	c.benchmark.SetProportion(util.Converter.StringToInteger(value, 100))
}
