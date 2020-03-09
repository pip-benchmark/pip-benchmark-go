package parameters

import (
	"strings"

	bench "github.com/pip-benchmark/pip-benchmark-go/benchmark"
	benchconf "github.com/pip-benchmark/pip-benchmark-go/runner/config"
)

type MeasurementTypeParameter struct {
	*bench.Parameter
	configuration *benchconf.ConfigurationManager
}

func NewMeasurementTypeParameter(configuration *benchconf.ConfigurationManager) *MeasurementTypeParameter {
	c := MeasurementTypeParameter{}
	c.Parameter = bench.NewParameter(
		"General.Benchmarking.MeasurementType",
		"Performance type: peak or nominal",
		"Peak",
	)
	c.configuration = configuration
	return &c
}

func (c *MeasurementTypeParameter) GetValue() string {
	if c.configuration.GetMeasurementType() == benchconf.Peak {
		return "Peak"
	}
	return "Nominal"
}

func (c *MeasurementTypeParameter) SetValue(value string) {
	value = strings.ToLower(value)
	if strings.HasPrefix(value, "p") {
		c.configuration.SetMeasurementType(benchconf.Peak)
		return
	}
	c.configuration.SetMeasurementType(benchconf.Nominal)
}
