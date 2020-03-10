package main

import (
	benchmark "github.com/pip-benchmark/pip-benchmark-go/benchmark"
)

var StandardHardwareBenchmarkSuite *tStandardHardwareBenchmarkSuite = newStandardHardwareBenchmarkSuite()

type tStandardHardwareBenchmarkSuite struct {
	*benchmark.BenchmarkSuite
	cpuBenchmarkTest   *tStandardCpuBenchmark
	diskBenchmarkTest  *tStandardDiskBenchmark
	videoBenchmarkTest *tStandardVideoBenchmark
}

func newStandardHardwareBenchmarkSuite() *tStandardHardwareBenchmarkSuite {
	c := tStandardHardwareBenchmarkSuite{}
	c.BenchmarkSuite = benchmark.NewBenchmarkSuite("StandardBenchmark", "Standard hardware benchmark")
	c.cpuBenchmarkTest = newStandardCpuBenchmark()
	c.AddBenchmark(c.cpuBenchmarkTest.Benchmark)

	c.diskBenchmarkTest = newStandardDiskBenchmark()
	c.AddBenchmark(c.diskBenchmarkTest.Benchmark)

	c.videoBenchmarkTest = newStandardVideoBenchmark()
	c.AddBenchmark(c.videoBenchmarkTest.Benchmark)
	return &c
}

func (c *tStandardHardwareBenchmarkSuite) CpuBenchmarkTest() *tStandardCpuBenchmark {
	return c.cpuBenchmarkTest
}

func (c *tStandardHardwareBenchmarkSuite) DiskBenchmarkTest() *tStandardDiskBenchmark {
	return c.diskBenchmarkTest
}

func (c *tStandardHardwareBenchmarkSuite) VideoBenchmarkTest() *tStandardVideoBenchmark {
	return c.videoBenchmarkTest
}
