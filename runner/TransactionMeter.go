package runner

import "time"

type TransactionMeter struct {
	*BenchmarkMeter
	transactionCounter int
}

func NewTransactionMeter() *TransactionMeter {
	c := TransactionMeter{}
	c.BenchmarkMeter = NewBenchmarkMeter()
	c.BenchmarkMeter.IPerfomedMesurement = &c
	c.Clear()
	return &c
}

func (c *TransactionMeter) IncrementTransactionCounter() {
	c.transactionCounter++
}

func (c *TransactionMeter) SetTransactionCounter(value int) {
	c.transactionCounter = value
}

func (c *TransactionMeter) PerformMeasurement() float64 {
	currentTime := time.Now()
	durationInMsecs := (currentTime.UnixNano() - c.LastMeasuredTime.UnixNano()) / int64(time.Millisecond)
	result := float64(c.transactionCounter) * 1000.0 / float64(durationInMsecs)
	c.LastMeasuredTime = currentTime
	c.transactionCounter = 0
	return result
}
