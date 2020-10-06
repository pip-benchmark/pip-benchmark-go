package runner

import "time"

type ResultAggregator struct {
	MaxErrorCount int

	results            *ResultsManager
	benchmarks         []*BenchmarkInstance
	transactionCounter int
	result             *BenchmarkResult
	transactionMeter   *TransactionMeter
	cpuLoadMeter       *CpuLoadMeter
	memoryUsageMeter   *MemoryUsageMeter
}

func NewResultAggregator(results *ResultsManager, benchmarks []*BenchmarkInstance) *ResultAggregator {
	c := ResultAggregator{}

	c.MaxErrorCount = 1000
	c.benchmarks = make([]*BenchmarkInstance, 0)
	c.transactionCounter = 0

	c.results = results
	c.benchmarks = benchmarks

	c.cpuLoadMeter = NewCpuLoadMeter()
	c.transactionMeter = NewTransactionMeter()
	c.memoryUsageMeter = NewMemoryUsageMeter()

	c.Start()
	return &c
}

func (c *ResultAggregator) Result() *BenchmarkResult {
	return c.result
}

func (c *ResultAggregator) Start() {
	c.result = NewBenchmarkResult()
	c.result.Benchmarks = c.benchmarks
	c.result.StartTime = time.Now()

	c.transactionCounter = 0
	c.transactionMeter.Clear()
	c.cpuLoadMeter.Clear()
	c.memoryUsageMeter.Clear()
}

func (c *ResultAggregator) IncrementCounter(increment int, now time.Time) {
	if now.IsZero() {
		now = time.Now()
	}
	c.transactionCounter += increment

	// If it's less then a second then wait
	measureInterval := time.Duration(now.UnixNano() - c.transactionMeter.LastMeasuredTime.UnixNano()) // / time.Millisecond
	if measureInterval >= time.Second && c.result != nil {
		// Perform measurements
		c.transactionMeter.SetTransactionCounter(c.transactionCounter)
		c.transactionCounter = 0
		c.transactionMeter.Measure()
		c.cpuLoadMeter.Measure()
		c.memoryUsageMeter.Measure()

		// Store measurement results
		c.result.ElapsedTime = time.Duration(now.UnixNano() - c.result.StartTime.UnixNano())
		c.result.PerformanceMeasurement = c.transactionMeter.Measurement()
		c.result.CpuLoadMeasurement = c.cpuLoadMeter.Measurement()
		c.result.MemoryUsageMeasurement = c.memoryUsageMeter.Measurement()

		c.results.NotifyUpdated(c.result)
	}
}

func (c *ResultAggregator) SendMessage(message string) {
	c.results.NotifyMessage(message)
}

func (c *ResultAggregator) ReportError(err error) {
	if len(c.result.Errors) < c.MaxErrorCount {
		c.result.Errors = append(c.result.Errors, err)
	}
	c.results.NotifyError(err)
}

func (c *ResultAggregator) Stop() {
	c.results.Add(c.result)
}
