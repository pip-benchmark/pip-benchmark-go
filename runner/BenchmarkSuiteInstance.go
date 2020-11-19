package runner

import (
	"sync"

	benchmark "github.com/pip-benchmark/pip-benchmark-go/benchmark"
)

type BenchmarkSuiteInstance struct {
	suite      *benchmark.BenchmarkSuite
	benchmarks []*BenchmarkInstance
}

func NewBenchmarkSuiteInstance(suite *benchmark.BenchmarkSuite) *BenchmarkSuiteInstance {
	c := BenchmarkSuiteInstance{}
	c.suite = suite
	c.benchmarks = make([]*BenchmarkInstance, 0)
	for _, benchmark := range c.suite.Benchmarks() {
		c.benchmarks = append(c.benchmarks, NewBenchmarkInstance(&c, benchmark))
	}
	return &c
}

func (c *BenchmarkSuiteInstance) Suite() *benchmark.BenchmarkSuite {
	return c.suite
}

func (c *BenchmarkSuiteInstance) Name() string {
	return c.suite.Name()
}

func (c *BenchmarkSuiteInstance) Description() string {
	return c.suite.Description()
}

func (c *BenchmarkSuiteInstance) Parameters() []*benchmark.Parameter {
	var result []*benchmark.Parameter = make([]*benchmark.Parameter, 0)
	parameters := c.suite.Parameters()
	for prop := range parameters {
		if parameter, ok := parameters[prop]; ok {

			//if param, convOk := parameter.(*benchmark.Parameter); convOk {
			result = append(result, parameter)
			//}

		}
	}
	return result
}

func (c *BenchmarkSuiteInstance) Benchmarks() []*BenchmarkInstance {
	return c.benchmarks
}

func (c *BenchmarkSuiteInstance) IsSelected() []*BenchmarkInstance {

	result := make([]*BenchmarkInstance, 0)

	for _, bench := range c.benchmarks {
		if bench.IsSelected() {
			result = append(result, bench)
		}
	}
	return result
}

func (c *BenchmarkSuiteInstance) SelectAll() {

	for i := range c.benchmarks {
		c.benchmarks[i].Select(true)
	}
}

func (c *BenchmarkSuiteInstance) SelectByName(benchmarkName string) {

	for i := range c.benchmarks {
		if c.benchmarks[i].Name() == benchmarkName {
			c.benchmarks[i].Select(true)
		}
	}
}

func (c *BenchmarkSuiteInstance) UnselectAll() {
	for i := range c.benchmarks {
		c.benchmarks[i].Select(false)
	}
}

func (c *BenchmarkSuiteInstance) UnselectByName(benchmarkName string) {
	for i := range c.benchmarks {
		if c.benchmarks[i].Name() == benchmarkName {
			c.benchmarks[i].Select(false)
		}
	}
}

func (c *BenchmarkSuiteInstance) SetUp(context benchmark.IExecutionContext) error {
	c.suite.SetContext(context)

	setErr := c.suite.IPrepared.SetUp()
	if setErr != nil {
		return setErr
	}

	var err error
	var wg sync.WaitGroup = sync.WaitGroup{}

	for _, benchmark := range c.benchmarks {
		wg.Add(1)
		go func(item *BenchmarkInstance) {
			defer wg.Done()
			if benchmark.IsSelected() {
				err = benchmark.SetUp(context)
			}
		}(benchmark)
	}
	wg.Wait()
	return err
}

func (c *BenchmarkSuiteInstance) TearDown() error {

	downErr := c.suite.IPrepared.TearDown()

	if downErr != nil {
		return downErr
	}

	var err error
	var wg sync.WaitGroup = sync.WaitGroup{}

	for _, benchmark := range c.benchmarks {
		wg.Add(1)
		go func(item *BenchmarkInstance) {
			defer wg.Done()
			if benchmark.IsSelected() {
				err = benchmark.TearDown()
			}
		}(benchmark)
	}
	wg.Wait()
	// if err != nil {
	c.suite.SetContext(nil)
	// }
	return err
}
