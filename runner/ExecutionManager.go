package runner

import (
	cerr "github.com/pip-services3-go/pip-services3-commons-go/errors"
)

type ExecutionManager struct {
	Configuration *ConfigurationManager
	Results       *ResultsManager

	updatedListeners []*ExecutionCallback
	running          bool
	strategy         *ExecutionStrategy
}

func NewExecutionManager(configuration *ConfigurationManager, results *ResultsManager) *ExecutionManager {
	c := ExecutionManager{}
	c.updatedListeners = make([]*ExecutionCallback, 0)
	c.running = false
	c.Configuration = configuration
	c.Results = results
	return &c
}

func (c *ExecutionManager) IsRunning() bool {
	return c.running
}

func (c *ExecutionManager) Start(benchmarks []*BenchmarkInstance) {
	c.Run(benchmarks, func(err error) {})
}

func (c *ExecutionManager) Run(benchmarks []*BenchmarkInstance, callback func(err error)) {

	if benchmarks == nil || len(benchmarks) == 0 {
		callback(cerr.NewError("There are no benchmarks to execute"))
		return
	}

	if c.running {
		c.Stop()
	}
	c.running = true

	c.Results.Clear()
	c.NotifyUpdated(Running)

	// Create requested execution strategy
	if c.Configuration.GetExecutionType() == Sequential {
		c.strategy = NewSequencialExecutionStrategy(c.Configuration, c.Results, c, benchmarks).ExecutionStrategy
	} else {
		c.strategy = NewProportionalExecutionStrategy(c.Configuration, c.Results, c, benchmarks).ExecutionStrategy
	}

	// Initialize parameters and start
	err := c.strategy.Start()
	if err != nil {
		c.Stop()
		if callback != nil {
			callback(err)
		}
	}
	callback(nil)
}

func (c *ExecutionManager) Stop() {
	if c.running {
		c.running = false

		if c.strategy != nil {
			c.strategy.Stop()
			c.strategy = nil
		}
		c.NotifyUpdated(Completed)
	}
}

func (c *ExecutionManager) AddUpdatedListener(listener *ExecutionCallback) {
	c.updatedListeners = append(c.updatedListeners, listener)
}

func (c *ExecutionManager) RemoveUpdatedListener(listener *ExecutionCallback) {
	for index := len(c.updatedListeners) - 1; index >= 0; index-- {
		if c.updatedListeners[index] == listener {
			if index == len(c.updatedListeners) {
				c.updatedListeners = c.updatedListeners[:index-1]
			} else {
				c.updatedListeners = append(c.updatedListeners[:index], c.updatedListeners[index+1:]...)
			}
		}
	}
}

func (c *ExecutionManager) NotifyUpdated(state ExecutionState) {
	for index := 0; index < len(c.updatedListeners); index++ {
		listener := c.updatedListeners[index]
		(*listener)(state)

	}
}
