package runner

import (
	"time"
)

type BenchmarkResult struct {
	Benchmarks             []*BenchmarkInstance
	StartTime              time.Time
	ElapsedTime            time.Duration
	PerformanceMeasurement *Measurement
	CpuLoadMeasurement     *Measurement
	MemoryUsageMeasurement *Measurement
	Errors                 []error
}

func NewBenchmarkResult() *BenchmarkResult {
	return &BenchmarkResult{
		Benchmarks:             make([]*BenchmarkInstance, 0),
		StartTime:              time.Now(),
		ElapsedTime:            0,
		PerformanceMeasurement: NewMeasurement(0, 0, 0, 0),
		CpuLoadMeasurement:     NewMeasurement(0, 0, 0, 0),
		MemoryUsageMeasurement: NewMeasurement(0, 0, 0, 0),
		Errors:                 make([]error, 0),
	}
}
