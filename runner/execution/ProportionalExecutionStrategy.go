package execution

import (
	"math/rand"
	"sync"
	"time"

	bench "github.com/pip-benchmark/pip-benchmark-go/runner/benchmarks"
	benchconf "github.com/pip-benchmark/pip-benchmark-go/runner/config"
	benchres "github.com/pip-benchmark/pip-benchmark-go/runner/results"
)

type ProportionalExecutionStrategy struct {
	*ExecutionStrategy
	running             bool
	aggregator          ResultAggregator
	ticksPerTransaction int64
	lastExecutedTime    time.Time
	stopTime            time.Time
	benchmarkCount      int
	onlyBenchmark       *bench.BenchmarkInstance
	timeout             chan bool
}

func NewProportionalExecutionStrategy(configuration *benchconf.ConfigurationManager, results *benchres.ResultsManager,
	execution interface{}, benchmarks []*bench.BenchmarkInstance) *ProportionalExecutionStrategy {
	c := ProportionalExecutionStrategy{}
	c.ExecutionStrategy = NewExecutionStrategy(configuration, results, execution, benchmarks)
	c.running = false
	c.ticksPerTransaction = 0
	c.ExecutionStrategy.IExecutionStrategy = &c
	c.aggregator = NewResultAggregator(results, benchmarks)
	return &c
}

func (c *ProportionalExecutionStrategy) Start() error {
	if c.running {
		return nil
	}

	c.running = true
	c.aggregator.Start()

	c.calculateProportionalRanges()

	if c.Configuration.GetMeasurementType() == benchconf.Nominal {
		c.ticksPerTransaction = 1000 / int64(c.Configuration.GetNominalRate())
	}

	// Initialize and start
	var wg sync.WaitGroup = sync.WaitGroup{}
	var err error
	for _, suite := range c.Suites {

		wg.Add(1)
		go func(s *bench.BenchmarkSuiteInstance) {
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

func (c *ProportionalExecutionStrategy) Stop() error {
	// Interrupt any wait
	if c.timeout != nil {
		c.timeout <- true
		c.timeout = nil
	}

	// Stop and cleanup execution
	if c.running {
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
			go func(s *bench.BenchmarkSuiteInstance) {
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

func (c *ProportionalExecutionStrategy) chooseBenchmarkProportionally() *bench.BenchmarkInstance {
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
		for {
			select {
			case <-ticker.C:
				c.lastExecutedTime = time.Now()
				callback(nil)
				ticker.Stop()
				return
			case <-clear:
				ticker.Stop()
				return
			}
		}
	}()
	c.timeout = clear
}

func (c *ProportionalExecutionStrategy) executeBenchmark(benchmark *bench.BenchmarkInstance, callback func(error)) {

	if benchmark == nil || benchmark.IsPassive() {
		// Delay if benchmarks are passive
		c.executeDelay(500, callback)
	} else {
		// Execute active benchmark
		err := benchmark.Execute()

		// Process force continue
		if err != nil && c.Configuration.ForceContinue {
			c.aggregator.ReportError(err)
			err = nil
		}

		// Increment counter
		now := time.Now()
		if err == nil {
			c.aggregator.IncrementCounter(1, now)
		}

		// Introduce delay to keep nominal rate
		if err == nil && c.Configuration.GetMeasurementType() == benchconf.Nominal {
			delay := c.ticksPerTransaction - (now.Unix() - c.lastExecutedTime.Unix())
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
	duration := c.Configuration.Duration
	if duration <= 0 {
		duration = 365 * 24 * 36000
	}
	c.stopTime = time.Now().Unix() + duration*1000

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
		for c.running && c.lastExecutedTime.Unix() < c.stopTime.Unix() {
			benchmark := c.onlyBenchmark
			if c.onlyBenchmark != nil {
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
		return c.Stop()
	}
	return nil
}
