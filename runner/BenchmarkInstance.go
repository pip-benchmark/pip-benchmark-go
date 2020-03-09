package runner

import (
	ix "github.com/adam-lavrik/go-imath/ix"
	benchmark "github.com/pip-benchmark/pip-benchmark-go/benchmark"
)

type BenchmarkInstance struct {
	suite      *BenchmarkSuiteInstance
	benchmark  *benchmark.Benchmark
	selected   bool
	proportion int
	startRange int
	endRange   int
}

func NewBenchmarkInstance(suite *BenchmarkSuiteInstance, benchmark *benchmark.Benchmark) *BenchmarkInstance {
	c := BenchmarkInstance{}
	c.suite = suite
	c.benchmark = benchmark
	c.selected = false
	c.proportion = 100
	return &c
}

func (c *BenchmarkInstance) Suite() *BenchmarkSuiteInstance {
	return c.suite
}

func (c *BenchmarkInstance) Benchmark() *benchmark.Benchmark {
	return c.benchmark
}

func (c *BenchmarkInstance) Name() string {
	return c.benchmark.Name()
}

func (c *BenchmarkInstance) FullName() string {
	return "" + c.suite.Name() + "." + c.Name()
}

func (c *BenchmarkInstance) Description() string {
	return c.benchmark.Description()
}

func (c *BenchmarkInstance) IsSelected() bool {
	return c.selected
}

func (c *BenchmarkInstance) Select(value bool) {
	c.selected = value
}

func (c *BenchmarkInstance) IsPassive() bool {
	return c.benchmark.Type() == "PassiveBenchmark"
}

func (c *BenchmarkInstance) GetProportion() int {
	return c.proportion
}

func (c *BenchmarkInstance) SetProportion(value int) {
	c.proportion = ix.Max(0, ix.Min(10000, value))
}

func (c *BenchmarkInstance) GetStartRange() int {
	return c.startRange
}

func (c *BenchmarkInstance) SetStartRange(value int) {
	c.startRange = value
}

func (c *BenchmarkInstance) GetendRange() int {
	return c.endRange
}

func (c *BenchmarkInstance) SetEndRange(value int) {
	c.endRange = value
}

func (c *BenchmarkInstance) WithinRange(proportion int) bool {
	return proportion >= c.startRange && proportion < c.endRange
}

func (c *BenchmarkInstance) SetUp(context benchmark.IExecutionContext) error {
	c.benchmark.SetContext(context)
	return c.benchmark.SetUp()
}

func (c *BenchmarkInstance) Execute() error {

	return c.benchmark.Execute()

}

func (c *BenchmarkInstance) TearDown() error {
	err := c.benchmark.TearDown()
	c.benchmark.SetContext(nil)
	return err
}
