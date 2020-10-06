package runner

import (
	"strings"

	bench "github.com/pip-benchmark/pip-benchmark-go/benchmark"
)

type MeasurementTypeParameter struct {
	*bench.Parameter
	configuration *ConfigurationManager
}

func NewMeasurementTypeParameter(configuration *ConfigurationManager) *MeasurementTypeParameter {
	c := MeasurementTypeParameter{}
	c.Parameter = bench.NewParameter(
		"General.Benchmarking.MeasurementType",
		"Performance type: peak or nominal",
		"Peak",
		"MeasurementTypeParameter",
	)
	c.configuration = configuration
	c.IParameter = &c
	return &c
}

func (c *MeasurementTypeParameter) GetValue() string {
	if c.configuration.GetMeasurementType() == Peak {
		return "Peak"
	}
	return "Nominal"
}

func (c *MeasurementTypeParameter) SetValue(value string) {
	value = strings.ToLower(value)
	if strings.HasPrefix(value, "p") {
		c.configuration.SetMeasurementType(Peak)
		return
	}
	c.configuration.SetMeasurementType(Nominal)
}
