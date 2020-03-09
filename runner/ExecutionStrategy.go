package runner

type ExecutionStrategy struct {
	IExecutionStrategy
	Configuration *ConfigurationManager
	Results       *ResultsManager
	Execution     IStrategyExecutor

	Benchmarks       []*BenchmarkInstance
	ActiveBenchmarks []*BenchmarkInstance
	Suites           []*BenchmarkSuiteInstance
}

func NewExecutionStrategy(configuration *ConfigurationManager,
	results *ResultsManager, execution IStrategyExecutor, benchmarks []*BenchmarkInstance) *ExecutionStrategy {
	c := ExecutionStrategy{}
	c.Configuration = configuration
	c.Results = results
	c.Execution = execution

	c.Benchmarks = benchmarks
	c.ActiveBenchmarks = c.getActiveBenchmarks(benchmarks)
	c.Suites = c.getBenchmarkSuites(benchmarks)
	return &c
}

func (c *ExecutionStrategy) getActiveBenchmarks(benchmarks []*BenchmarkInstance) []*BenchmarkInstance {

	result := make([]*BenchmarkInstance, 0)
	for _, bencmark := range benchmarks {
		if !bencmark.IsPassive() {
			result = append(result, bencmark)
		}
	}
	return result
}

func (c *ExecutionStrategy) getBenchmarkSuites(benchmarks []*BenchmarkInstance) []*BenchmarkSuiteInstance {
	suites := make([]*BenchmarkSuiteInstance, 0)
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
