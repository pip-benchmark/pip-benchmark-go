package runner

import (
	benchmark "github.com/pip-benchmark/pip-benchmark-go/benchmark"
)

type StandardBenchmarkSuite struct {
	*benchmark.BenchmarkSuite
	cpuBenchmark   *DefaultCpuBenchmark
	diskBenchmark  *DefaultDiskBenchmark
	videoBenchmark *DefaultVideoBenchmark
}

func NewStandardBenchmarkSuite() *StandardBenchmarkSuite {
	c := StandardBenchmarkSuite{}
	c.BenchmarkSuite = benchmark.NewBenchmarkSuite("StandardBenchmark", "Measures overall system performance")
	c.cpuBenchmark = NewDefaultCpuBenchmark()
	c.AddBenchmark(c.cpuBenchmark.Benchmark)

	c.diskBenchmark = NewDefaultDiskBenchmark()
	c.AddBenchmark(c.diskBenchmark.Benchmark)

	c.videoBenchmark = NewDefaultVideoBenchmark()
	c.AddBenchmark(c.videoBenchmark.Benchmark)

	c.CreateParameter("FilePath", "Path where test file is located on disk", "")
	c.CreateParameter("FileSize", "Size of the test file", "102400000")
	c.CreateParameter("ChunkSize", "Size of a chunk that read or writter from/to test file", "1024000")
	c.CreateParameter("OperationTypes", "Types of test operations: read, write or all", "all")
	return &c
}

func (c *StandardBenchmarkSuite) GetCpuBenchmark() *DefaultCpuBenchmark {
	return c.cpuBenchmark
}

func (c *StandardBenchmarkSuite) GetDiskBenchmark() *DefaultDiskBenchmark {
	return c.diskBenchmark
}

func (c *StandardBenchmarkSuite) GetVideoBenchmark() *DefaultVideoBenchmark {
	return c.videoBenchmark
}
