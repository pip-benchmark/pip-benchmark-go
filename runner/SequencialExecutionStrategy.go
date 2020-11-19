package runner

import (
	"sync"
	"time"

	cerr "github.com/pip-services3-go/pip-services3-commons-go/errors"
)

type SequencialExecutionStrategy struct {
	*ExecutionStrategy
	running bool
	current *ProportionalExecutionStrategy
	timeout chan bool
}

func NewSequencialExecutionStrategy(configuration *ConfigurationManager, results *ResultsManager,
	execution IStrategyExecutor, benchmarks []*BenchmarkInstance) *SequencialExecutionStrategy {
	c := SequencialExecutionStrategy{}
	c.running = false
	//super(configuration, results, execution, benchmarks);
	c.ExecutionStrategy = NewExecutionStrategy(configuration, results, execution, benchmarks)
	c.timeout = nil
	c.IExecutionStrategy = &c
	return &c
}

func (c *SequencialExecutionStrategy) IsStopped() bool {
	return !c.running
}

func (c *SequencialExecutionStrategy) Start() error {
	if c.Configuration.duration <= 0 {
		cerr.NewError("Duration was not set")
	}
	if c.running {
		return nil
	}
	c.running = true
	return c.execute()
}

func (c *SequencialExecutionStrategy) StopExecution() error {
	if c.timeout != nil {
		//clearTimeout(c.timeout);
		c.timeout <- true
		c.timeout = nil
	}

	if c.running == true {
		c.running = false

		if c.Execution != nil {
			c.Execution.Stop()
		}

		if c.current != nil {
			return c.current.StopExecution()
		} else {
			return nil
		}
	}
	return nil
}

func (c *SequencialExecutionStrategy) execute() error {

	var wg sync.WaitGroup = sync.WaitGroup{}
	var errGlobal error
	wg.Add(1)
	go func() {
		for _, benchmark := range c.Benchmarks {
			// Skip if benchmarking was interrupted
			if !c.running {
				continue
			}

			c.Results.NotifyMessage("Executing " + benchmark.Name() + " benchmark...")
			// Start embedded strategy
			c.current = NewProportionalExecutionStrategy(c.Configuration, c.Results, nil, []*BenchmarkInstance{benchmark})
			c.current.Start()

			// Wait for specified duration and stop embedded strategy
			ticker := time.NewTicker(time.Duration(c.Configuration.GetDuration()) * time.Second)
			c.timeout = make(chan bool)

			select {
			case <-ticker.C:
				c.timeout = nil
				ticker.Stop()
				err := c.current.StopExecution()
				if err != nil {
					c.current = nil
					errGlobal = err
				}
			case <-c.timeout:
				c.timeout = nil
				ticker.Stop()
			}

		}
		wg.Done()
	}()

	wg.Wait()
	if errGlobal != nil {
		return errGlobal
	}
	return c.StopExecution()
}
