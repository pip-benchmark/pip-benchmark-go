package parameters

import (
	bench "github.com/pip-benchmark/pip-benchmark-go/benchmark"
	benchconf "github.com/pip-benchmark/pip-benchmark-go/runner/config"
	benchconv "github.com/pip-benchmark/pip-benchmark-go/utilities"
)

type NominalRateParameter struct {
	*bench.Parameter
	configuration *benchconf.ConfigurationManager
}

func NewNominalRateParameter(configuration *benchconf.ConfigurationManager) *NominalRateParameter {
	c := NominalRateParameter{}
	c.Parameter = bench.NewParameter(
		"General.Benchmarking.NominalRate",
		"Rate for nominal benchmarking in TPS",
		"1",
	)
	c.configuration = configuration
	return &c
}

func (c *NominalRateParameter) GetValue() string {
	return benchconv.Converter.DoubleToString(c.configuration.GetNominalRate())
}

func (c *NominalRateParameter) SetValue(value string) {
	c.configuration.SetNominalRate(benchconv.Converter.StringToDouble(value, 1.0))
}
