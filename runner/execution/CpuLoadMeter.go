package execution

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
)

type CpuLoadMeter struct {
	*BenchmarkMeter
	lastTotalIdle float64
	lastTotal     float64
}

func NewCpuLoadMeter() *CpuLoadMeter {
	c := CpuLoadMeter{}
	c.BenchmarkMeter = NewBenchmarkMeter()
	c.BenchmarkMeter.IPerfomedMesurement = &c
	return &c
}

func (c *CpuLoadMeter) Clear() {
	c.lastTotalIdle = 0
	c.lastTotal = 0
	c.BenchmarkMeter.Clear()
}

func (c *CpuLoadMeter) PerformMeasurement() float64 {
	// Initialize current values
	currentTime := time.Now()
	currentTotalIdle := 0.0
	currentTotal := 0.0

	// Calculate current values
	cpus, cpuErr := cpu.Times(true)
	if cpuErr != nil {

	}
	cpuCount := len(cpus)
	for index := 0; index < cpuCount; index++ {
		cpu := cpus[index]
		currentTotal += cpu.Total() - cpu.Idle
		currentTotalIdle += cpu.Idle
	}
	currentTotal = currentTotal / (float64)(cpuCount)
	currentTotalIdle = currentTotalIdle / (float64)(cpuCount)

	// Calculate CPU usage
	result := 0.0
	if !c.LastMeasuredTime.IsZero() {
		elapsed := currentTime.UnixNano() - c.LastMeasuredTime.UnixNano()
		// Calculate only for 100 ms or more
		if time.Duration(elapsed) > 100*time.Millisecond {
			totalDifference := currentTotal - c.lastTotal
			idleDifference := currentTotalIdle - c.lastTotalIdle
			result = 100.0 - float64(100*idleDifference/totalDifference)
		}
	}

	// Save current values as last values
	c.LastMeasuredTime = currentTime
	c.lastTotalIdle = currentTotalIdle
	c.lastTotal = currentTotal

	return result
}
