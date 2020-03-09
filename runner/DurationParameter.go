package runner

import (
	bench "github.com/pip-benchmark/pip-benchmark-go/benchmark"
	util "github.com/pip-benchmark/pip-benchmark-go/utilities"
)

type DurationParameter struct {
	*bench.Parameter
	configuration *ConfigurationManager
}

func NewDurationParameter(configuration *ConfigurationManager) *DurationParameter {
	c := DurationParameter{}
	c.Parameter = bench.NewParameter(
		"General.Benchmarking.Duration",
		"Duration of benchmark execution in seconds",
		"60",
		"DurationParameter",
	)
	c.configuration = configuration
	return &c
}

func (c *DurationParameter) GetValue() string {
	return util.Converter.IntegerToString(int(c.configuration.GetDuration()))
}

func (c *DurationParameter) SetValue(value string) {
	c.configuration.SetDuration(int64(util.Converter.StringToInteger(value, 60)))
}
