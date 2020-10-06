package runner

import (
	"time"

	benchmark "github.com/pip-benchmark/pip-benchmark-go/benchmark"
)

//implements IExecutionContext
type ExecutionContext struct {
	suite      *BenchmarkSuiteInstance
	aggregator *ResultAggregator
	strategy   *ExecutionStrategy
}

func NewExecutionContext(suite *BenchmarkSuiteInstance,
	aggregator *ResultAggregator, strategy *ExecutionStrategy) *ExecutionContext {

	c := ExecutionContext{}
	c.aggregator = aggregator
	c.suite = suite
	c.strategy = strategy
	return &c
}

func (c *ExecutionContext) GetParameters() map[string]*benchmark.Parameter {
	return c.suite.Suite().Parameters()
}

func (c *ExecutionContext) IncrementCounter(increment int) {
	c.aggregator.IncrementCounter(increment, time.Now())
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
	c.strategy.StopExecution()
}
