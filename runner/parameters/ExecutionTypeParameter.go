package parameters

import (
	"strings"

	bench "github.com/pip-benchmark/pip-benchmark-go/benchmark"
	benchconf "github.com/pip-benchmark/pip-benchmark-go/runner/config"
)

type ExecutionTypeParameter struct {
	*bench.Parameter
	configuration *benchconf.ConfigurationManager
}

func NewExecutionTypeParameter(configuration *benchconf.ConfigurationManager) *ExecutionTypeParameter {
	c := ExecutionTypeParameter{}
	c.Parameter = bench.NewParameter(
		"General.Benchmarking.ExecutionType",
		"Execution type: proportional or sequencial",
		"Proportional",
	)
	c.configuration = configuration
	return &c
}

func (c *ExecutionTypeParameter) GetValue() string {

	if c.configuration.GetExecutionType() == benchconf.Proportional {
		return "Proportional"
	}
	return "Sequencial"
}

func (c *ExecutionTypeParameter) SetValue(value string) {
	value = strings.ToLower(value)
	if strings.HasPrefix(value, "p") {
		c.configuration.SetExecutionType(benchconf.Proportional)
		return
	}
	c.configuration.SetExecutionType(benchconf.Sequential)
}
