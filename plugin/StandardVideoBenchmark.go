package main

import (
	benchmark "github.com/pip-benchmark/pip-benchmark-go/benchmark"
)

var StandardVideoBenchmark *tStandardVideoBenchmark = newStandardVideoBenchmark()

type tStandardVideoBenchmark struct {
	*benchmark.PassiveBenchmark
}

func newStandardVideoBenchmark() *tStandardVideoBenchmark {
	c := tStandardVideoBenchmark{
		PassiveBenchmark: benchmark.NewPassiveBenchmark("Video", "Measures speed of drawing graphical primitives"),
	}
	return &c
}
