package runner

import (
	"strconv"

	benchmark "github.com/pip-benchmark/pip-benchmark-go/benchmark"
)

type BenchmarkBuilder struct {
	Runner *BenchmarkRunner
}

func NewBenchmarkBuilder() *BenchmarkBuilder {
	c := BenchmarkBuilder{}
	c.Runner = NewBenchmarkRunner()
	return &c
}

func (c *BenchmarkBuilder) ForceContinue(isForceContinue bool) *BenchmarkBuilder {
	c.Runner.configuration.forceContinue = isForceContinue
	return c
}

func (c *BenchmarkBuilder) MeasureAs(measurementType MeasurementType) *BenchmarkBuilder {
	c.Runner.configuration.measurementType = measurementType
	return c
}

func (c *BenchmarkBuilder) WithNominalRate(nominalRate float64) *BenchmarkBuilder {
	c.Runner.configuration.nominalRate = nominalRate
	return c
}

func (c *BenchmarkBuilder) ExecuteAs(executionType ExecutionType) *BenchmarkBuilder {
	c.Runner.configuration.executionType = executionType
	return c
}

func (c *BenchmarkBuilder) ForDuration(duration int64) *BenchmarkBuilder {
	c.Runner.configuration.duration = duration
	return c
}

func (c *BenchmarkBuilder) AddSuite(suite *benchmark.BenchmarkSuite) *BenchmarkBuilder {
	c.Runner.benchmarks.AddSuite(suite)
	return c
}

func (c *BenchmarkBuilder) WithParameter(name string, value string) *BenchmarkBuilder {
	parameters := make(map[string]string)
	parameters[name] = value
	c.Runner.parameters.Set(parameters)
	return c
}

func (c *BenchmarkBuilder) WithBenchmark(name string) *BenchmarkBuilder {
	c.Runner.benchmarks.SelectByName([]string{name})
	return c
}

func (c *BenchmarkBuilder) WithProportionalBenchmark(name string, proportion float32) *BenchmarkBuilder {
	c.Runner.benchmarks.SelectByName([]string{name})
	c.WithParameter(name+".Proportion", strconv.FormatFloat((float64)(proportion), 'f', 3, 32))
	return c
}

func (c *BenchmarkBuilder) WithAllBenchmarks() *BenchmarkBuilder {
	c.Runner.benchmarks.SelectAll()
	return c
}

func (c *BenchmarkBuilder) Create() *BenchmarkRunner {
	result := c.Runner
	c.Runner = NewBenchmarkRunner()
	return result
}
