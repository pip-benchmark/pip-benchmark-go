package standartbenchmarks

import (
	benchmark "github.com/pip-benchmark/pip-benchmark-go/benchmark"
)

type StandardVideoBenchmark struct {
	*benchmark.PassiveBenchmark
}

func NewStandardVideoBenchmark() *StandardVideoBenchmark {
	c := StandardVideoBenchmark{
		PassiveBenchmark: benchmark.NewPassiveBenchmark("Video", "Measures speed of drawing graphical primitives"),
	}
	return &c
}
