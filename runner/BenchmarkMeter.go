package runner

import (
	"math"
	"time"
)

type BenchmarkMeter struct {
	LastMeasuredTime     time.Time
	currentValue         float64
	minValue             float64
	maxValue             float64
	averageValue         float64
	sumOfValues          float64
	numberOfMeasurements int
	IPerfomedMesurement
}

func NewBenchmarkMeter() *BenchmarkMeter {
	c := BenchmarkMeter{}
	c.Clear()
	return &c
}

func (c *BenchmarkMeter) Measurement() *Measurement {
	return NewMeasurement(c.currentValue, c.minValue,
		c.averageValue, c.maxValue)
}

func (c *BenchmarkMeter) GetLastMeasuredTime() time.Time {
	return c.LastMeasuredTime
}

func (c *BenchmarkMeter) SetLastMeasuredTime(value time.Time) {
	c.LastMeasuredTime = value
}

func (c *BenchmarkMeter) GetCurrentValue() float64 {
	return c.currentValue
}

func (c *BenchmarkMeter) SetCurrentValue(value float64) {
	c.currentValue = value
}

func (c *BenchmarkMeter) GetMinValue() float64 {
	if c.minValue < math.MaxFloat64 {
		return c.minValue
	}
	return 0.0
}

func (c *BenchmarkMeter) SetMinValue(value float64) {
	c.minValue = value
}

func (c *BenchmarkMeter) GetmaxValue() float64 {
	if c.maxValue > math.SmallestNonzeroFloat64 {
		return c.maxValue
	}
	return 0.0
}

func (c *BenchmarkMeter) SetMaxValue(value float64) {
	c.minValue = value
}

func (c *BenchmarkMeter) GetAverageValue() float64 {
	return c.averageValue
}

func (c *BenchmarkMeter) SetAverageValue(value float64) {
	c.averageValue = value
}

func (c *BenchmarkMeter) Clear() {
	c.LastMeasuredTime = time.Now()
	//c.currentValue = c.PerformMeasurement()
	c.currentValue = 0.0
	c.minValue = math.MaxFloat64
	c.maxValue = math.SmallestNonzeroFloat64
	c.averageValue = 0
	c.sumOfValues = 0
	c.numberOfMeasurements = 0
}

func (c *BenchmarkMeter) CalculateAggregates() {
	c.sumOfValues += c.currentValue
	c.numberOfMeasurements++
	c.averageValue = c.sumOfValues / (float64)(c.numberOfMeasurements)
	c.maxValue = math.Max(c.maxValue, c.currentValue)
	c.minValue = math.Min(c.minValue, c.currentValue)
}

func (c *BenchmarkMeter) Measure() float64 {
	c.currentValue = c.PerformMeasurement()
	c.LastMeasuredTime = time.Now()
	c.CalculateAggregates()
	return c.currentValue
}
