package runner

import (
	bench "github.com/pip-benchmark/pip-benchmark-go/benchmark"
	util "github.com/pip-benchmark/pip-benchmark-go/utilities"
)

type NominalRateParameter struct {
	*bench.Parameter
	configuration *ConfigurationManager
}

func NewNominalRateParameter(configuration *ConfigurationManager) *NominalRateParameter {
	c := NominalRateParameter{}
	c.Parameter = bench.NewParameter(
		"General.Benchmarking.NominalRate",
		"Rate for nominal benchmarking in TPS",
		"1",
		"NominalRateParameter",
	)
	c.configuration = configuration
	c.IParameter = &c
	return &c
}

func (c *NominalRateParameter) GetValue() string {
	return util.Converter.DoubleToString(c.configuration.GetNominalRate())
}

func (c *NominalRateParameter) SetValue(value string) {
	c.configuration.SetNominalRate(util.Converter.StringToDouble(value, 1.0))
}
