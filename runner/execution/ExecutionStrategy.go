package execution

import (
	bench "github.com/pip-benchmark/pip-benchmark-go/runner/benchmarks"
	benchconfig "github.com/pip-benchmark/pip-benchmark-go/runner/config"
	benchresult "github.com/pip-benchmark/pip-benchmark-go/runner/results"
)

type ExecutionStrategy struct {
	IExecutionStrategy
	Configuration *benchconfig.ConfigurationManager
	Results       *benchresult.ResultsManager
	Execution     interface{}

	Benchmarks       []*bench.BenchmarkInstance
	ActiveBenchmarks []*bench.BenchmarkInstance
	Suites           []*bench.BenchmarkSuiteInstance
}

func NewExecutionStrategy(configuration *benchconfig.ConfigurationManager,
	results *benchresult.ResultsManager, execution interface{}, benchmarks []*bench.BenchmarkInstance) *ExecutionStrategy {
	c := ExecutionStrategy{}
	c.Configuration = configuration
	c.Results = results
	c.Execution = execution

	c.Benchmarks = benchmarks
	c.ActiveBenchmarks = c.getActiveBenchmarks(benchmarks)
	c.Suites = c.getBenchmarkSuites(benchmarks)
	return &c
}

func (c *ExecutionStrategy) getActiveBenchmarks(benchmarks []*bench.BenchmarkInstance) []*bench.BenchmarkInstance {

	result := make([]*bench.BenchmarkInstance, 0)
	for _, bencmark := range benchmarks {
		if !bencmark.IsPassive() {
			result = append(result, bencmark)
		}
	}
	return result
}

func (c *ExecutionStrategy) getBenchmarkSuites(benchmarks []*bench.BenchmarkInstance) []*bench.BenchmarkSuiteInstance {
	suites := make([]*bench.BenchmarkSuiteInstance, 0)
	var exists bool = false
	for index := 0; index < len(benchmarks); index++ {
		benchmark := benchmarks[index]
		suite := benchmark.Suite()
		exists = false
		for i := range suites {
			if suite == suites[i] {
				exists = true
				break
			}
		}
		if !exists {
			suites = append(suites, suite)
		}
	}
	return suites
}

// public abstract isStopped bool
// public abstract start() error
// public abstract stop() error
