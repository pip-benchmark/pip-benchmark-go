package standardbenchmarks

import (
	benchmark "github.com/pip-benchmark/pip-benchmark-go/benchmark"
)

type StandardHardwareBenchmarkSuite struct {
	*benchmark.BenchmarkSuite
	cpuBenchmarkTest   *StandardCpuBenchmark
	diskBenchmarkTest  *StandardDiskBenchmark
	videoBenchmarkTest *StandardVideoBenchmark
}

func NewStandardHardwareBenchmarkSuite() *StandardHardwareBenchmarkSuite {
	c := StandardHardwareBenchmarkSuite{}
	c.BenchmarkSuite = benchmark.NewBenchmarkSuite("StandardBenchmark", "Standard hardware benchmark")
	c.cpuBenchmarkTest = NewStandardCpuBenchmark()
	c.AddBenchmark(c.cpuBenchmarkTest.Benchmark)

	c.diskBenchmarkTest = NewStandardDiskBenchmark()
	c.AddBenchmark(c.diskBenchmarkTest.Benchmark)

	c.videoBenchmarkTest = NewStandardVideoBenchmark()
	c.AddBenchmark(c.videoBenchmarkTest.Benchmark)
	return &c
}

func (c *StandardHardwareBenchmarkSuite) CpuBenchmarkTest() *StandardCpuBenchmark {
	return c.cpuBenchmarkTest
}

func (c *StandardHardwareBenchmarkSuite) DiskBenchmarkTest() *StandardDiskBenchmark {
	return c.diskBenchmarkTest
}

func (c *StandardHardwareBenchmarkSuite) VideoBenchmarkTest() *StandardVideoBenchmark {
	return c.videoBenchmarkTest
}
