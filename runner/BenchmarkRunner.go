package runner

type BenchmarkRunner struct {
	configuration *ConfigurationManager
	results       *ResultsManager
	parameters    *ParametersManager
	benchmarks    *BenchmarksManager
	execution     *ExecutionManager
	report        *ReportGenerator
	environment   *EnvironmentManager
}

func NewBenchmarkRunner() *BenchmarkRunner {
	c := BenchmarkRunner{}
	c.configuration = NewConfigurationManager()
	c.results = NewResultsManager()
	c.parameters = NewParametersManager(c.configuration)
	c.benchmarks = NewBenchmarksManager(c.parameters)
	c.execution = NewExecutionManager(c.configuration, c.results)
	c.environment = NewEnvironmentManager()
	c.report = NewReportGenerator(c.configuration, c.results,
		c.parameters, c.benchmarks, c.environment)
	return &c
}

func (c *BenchmarkRunner) Configuration() *ConfigurationManager {
	return c.configuration
}

func (c *BenchmarkRunner) Results() *ResultsManager {
	return c.results
}

func (c *BenchmarkRunner) Parameters() *ParametersManager {
	return c.parameters
}

func (c *BenchmarkRunner) Execution() *ExecutionManager {
	return c.execution
}

func (c *BenchmarkRunner) Benchmarks() *BenchmarksManager {
	return c.benchmarks
}

func (c *BenchmarkRunner) Report() *ReportGenerator {
	return c.report
}

func (c *BenchmarkRunner) Environment() *EnvironmentManager {
	return c.environment
}

func (c *BenchmarkRunner) IsRunning() bool {
	return c.execution.IsRunning()
}

func (c *BenchmarkRunner) Start() {
	c.execution.Start(c.benchmarks.IsSelected())
}

func (c *BenchmarkRunner) Stop() {
	c.execution.Stop()
}

func (c *BenchmarkRunner) Run(callback func(error)) {
	c.execution.Run(c.benchmarks.IsSelected(), callback)
}
