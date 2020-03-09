package execution

import (
	"runtime"
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
	stat := runtime.MemStats{}
	runtime.ReadMemStats(&stat)
	return float64((stat.Sys - stat.Frees) / 1024 / 1024)
}
