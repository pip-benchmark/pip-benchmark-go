package environment

import (
	"math"

	"github.com/pip-benchmark/pip-benchmark-go/benchmark"
)

type DefaultCpuBenchmark struct {
	*benchmark.Benchmark
	numberOfAttempts int
}

func NewDefaultCpuBenchmark() *DefaultCpuBenchmark {
	c := DefaultCpuBenchmark{}
	c.Benchmark = benchmark.NewBenchmark("CPU", "Measures CPU performance")
	c.Benchmark.IExecutable = &c
	c.numberOfAttempts = 20000
	return &c
}

func (c *DefaultCpuBenchmark) Execute() error {
	// Count increment, comparison and goto for 1 arithmetic operation
	for value := float64(0); value < float64(c.numberOfAttempts); value++ {
		// #1
		result1 := value + value
		result2 := result1 - value
		result3 := result1 * result2
		result4 := result2 / result3
		math.Log(result4)

		// #2
		result1 = value + value
		result2 = result1 - value
		result3 = result1 * result2
		result4 = result2 / result3
		math.Log(result4)

		// #3
		result1 = value + value
		result2 = result1 - value
		result3 = result1 * result2
		result4 = result2 / result3
		math.Log(result4)

		// #4
		result1 = value + value
		result2 = result1 - value
		result3 = result1 * result2
		result4 = result2 / result3
		math.Log(result4)

		// #5
		result1 = value + value
		result2 = result1 - value
		result3 = result1 * result2
		result4 = result2 / result3
		math.Log(result4)

		// #6
		result1 = value + value
		result2 = result1 - value
		result3 = result1 * result2
		result4 = result2 / result3
		math.Log(result4)

		// #7
		result1 = value + value
		result2 = result1 - value
		result3 = result1 * result2
		result4 = result2 / result3
		math.Log(result4)

		// #8
		result1 = value + value
		result2 = result1 - value
		result3 = result1 * result2
		result4 = result2 / result3
		math.Log(result4)

		// #9
		result1 = value + value
		result2 = result1 - value
		result3 = result1 * result2
		result4 = result2 / result3
		math.Log(result4)

		// #10
		result1 = value + value
		result2 = result1 - value
		result3 = result1 * result2
		result4 = result2 / result3
		math.Log(result4)
	}
	return nil
}
