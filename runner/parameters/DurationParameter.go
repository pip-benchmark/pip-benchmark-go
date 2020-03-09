package parameters

import (
	bench "github.com/pip-benchmark/pip-benchmark-go/benchmark"
	runnerconf "github.com/pip-benchmark/pip-benchmark-go/runner/config"
	benchconv "github.com/pip-benchmark/pip-benchmark-go/utilities"
)

type DurationParameter struct {
	*bench.Parameter
	configuration *runnerconf.ConfigurationManager
}

func NewDurationParameter(configuration *runnerconf.ConfigurationManager) *DurationParameter {
	c := DurationParameter{}
	c.Parameter = bench.NewParameter(
		"General.Benchmarking.Duration",
		"Duration of benchmark execution in seconds",
		"60",
	)
	c.configuration = configuration
	return &c
}

func (c *DurationParameter) GetValue() string {
	return benchconv.Converter.IntegerToString(int(c.configuration.GetDuration()))
}

func (c *DurationParameter) SetValue(value string) {
	c.configuration.SetDuration(int64(benchconv.Converter.StringToInteger(value, 60)))
}
