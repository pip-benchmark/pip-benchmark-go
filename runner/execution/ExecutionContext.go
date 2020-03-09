package execution

import (
	bench "github.com/pip-benchmark/pip-benchmark-go/runner/benchmarks"
)

//implements IExecutionContext
type ExecutionContext struct {
	suite      *bench.BenchmarkSuiteInstance
	aggregator *ResultAggregator
	strategy   *ExecutionStrategy
}

func NewExecutionContext(suite *bench.BenchmarkSuiteInstance,
	aggregator *ResultAggregator, strategy *ExecutionStrategy) *ExecutionContext {

	c := ExecutionContext{}
	c.aggregator = aggregator
	c.suite = suite
	c.strategy = strategy
	return &c
}

func (c *ExecutionContext) GetParameters() map[string]interface{} {
	return c.suite.Suite().Parameters()
}

func (c *ExecutionContext) IncrementCounter(increment int) {
	c.aggregator.IncrementCounter(increment)
}

func (c *ExecutionContext) SendMessage(message string) {
	c.aggregator.SendMessage(message)
}

func (c *ExecutionContext) ReportError(err error) {
	c.aggregator.ReportError(err)
}

func (c *ExecutionContext) IsStopped() bool {
	return c.strategy.IsStopped()
}

func (c *ExecutionContext) Stop() {
	c.strategy.Stop()
}
