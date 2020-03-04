package standartbenchmarks

import (
	"math/rand"
	"time"

	benchmark "github.com/pip-benchmark/pip-benchmark-go/benchmark"
)

type UtilityBenchmarkSuite struct {
	*benchmark.BenchmarkSuite
}

func NewUtilityBenchmarkSuite() *UtilityBenchmarkSuite {

	c := UtilityBenchmarkSuite{}
	c.BenchmarkSuite = benchmark.NewBenchmarkSuite("Utility", "Set of utility benchmark tests")
	c.CreateBenchmark("Empty", "Does nothing", c.ExecuteEmpty)
	c.CreateBenchmark("RandomDelay", "Introduces random delay to measuring thread", c.executeRandomDelay)
	return &c
}

func (c *UtilityBenchmarkSuite) ExecuteEmpty() error {
	// This is an empty benchmark
	return nil
}

func (c *UtilityBenchmarkSuite) executeRandomDelay() error {
	select {
	case <-time.After(time.Duration(rand.Intn(1000)) * time.Millisecond):
	}

	return nil
}
