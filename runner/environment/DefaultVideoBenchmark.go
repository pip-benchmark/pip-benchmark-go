package environment

import (
	benchmark "github.com/pip-benchmark/pip-benchmark-go/benchmark"
)

type DefaultVideoBenchmark struct {
	*benchmark.PassiveBenchmark
}

func NewDefaultVideoBenchmark() *DefaultVideoBenchmark {
	c := DefaultVideoBenchmark{
		PassiveBenchmark: benchmark.NewPassiveBenchmark("Video", "Measures speed of drawing graphical primitives"),
	}
	return &c
}
