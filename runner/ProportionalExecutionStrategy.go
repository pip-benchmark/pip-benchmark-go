package runner

import (
	"math/rand"
	"sync"
	"time"
)

type ProportionalExecutionStrategy struct {
	*ExecutionStrategy
	running             bool
	aggregator          *ResultAggregator
	ticksPerTransaction int64
	lastExecutedTime    time.Time
	stopTime            time.Time
	benchmarkCount      int
	onlyBenchmark       *BenchmarkInstance
	timeout             chan bool
}

func NewProportionalExecutionStrategy(configuration *ConfigurationManager, results *ResultsManager,
	execution IStrategyExecutor, benchmarks []*BenchmarkInstance) *ProportionalExecutionStrategy {
	c := ProportionalExecutionStrategy{}
	c.ExecutionStrategy = NewExecutionStrategy(configuration, results, execution, benchmarks)
	c.running = false
	c.ticksPerTransaction = 0
	c.IExecutionStrategy = &c
	c.aggregator = NewResultAggregator(results, benchmarks)
	c.timeout = nil
	return &c
}

func (c *ProportionalExecutionStrategy) Start() error {
	if c.running {
		return nil
	}

	c.running = true
	c.aggregator.Start()

	c.calculateProportionalRanges()

	if c.Configuration.GetMeasurementType() == Nominal {
		c.ticksPerTransaction = 1000 / int64(c.Configuration.GetNominalRate())
	}

	// Initialize and start
	var wg sync.WaitGroup = sync.WaitGroup{}
	var err error
	for _, suite := range c.Suites {

		wg.Add(1)
		go func(s *BenchmarkSuiteInstance) {
			defer wg.Done()
			context := NewExecutionContext(s, c.aggregator, c.ExecutionStrategy)
			setupErr := s.SetUp(context)
			if setupErr != nil {
				err = setupErr
			}
		}(suite)

	}

	wg.Wait()

	// Abort if initialization failed
	if err != nil {
		c.aggregator.ReportError(err)
		return err
	}

	// Execute benchmarks
	return c.execute()
}

func (c *ProportionalExecutionStrategy) IsStopped() bool {
	return !c.running
}

func (c *ProportionalExecutionStrategy) StopExecution() error {
	// Interrupt any wait
	if c.timeout != nil {
		c.timeout <- true
		c.timeout = nil
	}

	// Stop and cleanup execution
	if c.running == true {
		c.running = false
		c.aggregator.Stop()

		if c.Execution != nil {
			c.Execution.Stop()
		}

		// Deinitialize tests
		var wg sync.WaitGroup = sync.WaitGroup{}
		var err error
		for _, suite := range c.Suites {

			wg.Add(1)
			go func(s *BenchmarkSuiteInstance) {
				defer wg.Done()
				teardownErr := s.TearDown()
				if teardownErr != nil {
					err = teardownErr
				}
			}(suite)
		}
		wg.Wait()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *ProportionalExecutionStrategy) calculateProportionalRanges() {
	totalProportion := 0
	for _, benchmark := range c.ActiveBenchmarks {
		totalProportion += benchmark.GetProportion()
	}

	startRange := 0
	for _, benchmark := range c.ActiveBenchmarks {
		normalizedProportion := benchmark.GetProportion() / totalProportion
		benchmark.SetStartRange(startRange)
		benchmark.SetEndRange(startRange + normalizedProportion)
		startRange += normalizedProportion
	}
}

func (c *ProportionalExecutionStrategy) chooseBenchmarkProportionally() *BenchmarkInstance {
	proportion := rand.Int()

	for _, benchmark := range c.ActiveBenchmarks {
		if benchmark.WithinRange(proportion) {
			return benchmark
		}
	}
	return nil
}

func (c *ProportionalExecutionStrategy) executeDelay(delay int, callback func(error)) {

	interval := time.Duration(delay) * time.Millisecond

	ticker := time.NewTicker(interval)
	clear := make(chan bool)

	go func() {
		select {
		case <-ticker.C:
			c.lastExecutedTime = time.Now()
			c.timeout = nil
			ticker.Stop()
			callback(nil)
			return
		case <-clear:
			ticker.Stop()
			c.timeout = nil
			return
		}
	}()
	c.timeout = clear
}

func (c *ProportionalExecutionStrategy) executeBenchmark(benchmark *BenchmarkInstance, callback func(error)) {

	if benchmark == nil || benchmark.IsPassive() {
		// Delay if benchmarks are passive
		c.executeDelay(500, callback)
	} else {
		// Execute active benchmark
		err := benchmark.Execute()

		// Process force continue
		if err != nil && c.Configuration.GetForceContinue() {
			c.aggregator.ReportError(err)
			err = nil
		}

		// Increment counter
		now := time.Now()
		if err == nil {
			c.aggregator.IncrementCounter(1, now)
		}

		// Introduce delay to keep nominal rate
		if err == nil && c.Configuration.GetMeasurementType() == Nominal {
			delay := c.ticksPerTransaction - (now.UnixNano()-c.lastExecutedTime.UnixNano())/int64(time.Millisecond)
			c.lastExecutedTime = now

			if delay > 0 {
				c.executeDelay(int(delay), callback)
			} else {
				callback(err)
			}
		} else {
			c.lastExecutedTime = now
			callback(err)
		}
	}
	// // Process force continue
	// if c.Configuration.forceContinue {
	// 	c.aggregator.ReportError(ex)
	// 	callback(nil)
	// } else {
	// 	callback(ex)
	// }
}

func (c *ProportionalExecutionStrategy) execute() error {
	c.lastExecutedTime = time.Now()
	duration := c.Configuration.GetDuration()
	if duration <= 0 {
		duration = 365 * 24 * 36000
	}

	c.stopTime = time.Now().Add(time.Duration(duration) * time.Second) // time.Milisecon

	c.benchmarkCount = len(c.Benchmarks)
	c.onlyBenchmark = nil
	if c.benchmarkCount == 1 {
		c.onlyBenchmark = c.Benchmarks[0]
	}

	// Execute benchmarks
	var errGlobal error
	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for c.running && c.lastExecutedTime.UnixNano() < c.stopTime.UnixNano() {
			benchmark := c.onlyBenchmark
			if c.onlyBenchmark == nil {
				benchmark = c.chooseBenchmarkProportionally()
			}
			//called := 0
			var wg2 sync.WaitGroup = sync.WaitGroup{}
			wg2.Add(1)
			c.executeBenchmark(
				benchmark,
				func(err error) {
					//process.nextTick(callback, err)
					if err != nil {
						errGlobal = err
					}
					wg2.Done()
				})
			wg2.Wait()
		}
	}()
	wg.Wait()
	if errGlobal != nil {
		return c.StopExecution()
	}
	return nil
}
