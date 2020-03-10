package runner

import (
	mem "github.com/mackerelio/go-osstat/memory"
)

type MemoryUsageMeter struct {
	*BenchmarkMeter
}

func NewMemoryUsageMeter() *MemoryUsageMeter {
	c := MemoryUsageMeter{}
	c.BenchmarkMeter = NewBenchmarkMeter()
	c.BenchmarkMeter.IPerfomedMesurement = &c
	return &c
}

func (c *MemoryUsageMeter) PerformMeasurement() float64 {
	memory, err := mem.Get()
	if err != nil {
		return 0.0
	}

	return float64((memory.Total - memory.Free) / 1024 / 1024)
}
