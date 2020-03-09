package runner

import (
	"strings"

	bench "github.com/pip-benchmark/pip-benchmark-go/benchmark"
)

type ExecutionTypeParameter struct {
	*bench.Parameter
	configuration *ConfigurationManager
}

func NewExecutionTypeParameter(configuration *ConfigurationManager) *ExecutionTypeParameter {
	c := ExecutionTypeParameter{}
	c.Parameter = bench.NewParameter(
		"General.Benchmarking.ExecutionType",
		"Execution type: proportional or sequencial",
		"Proportional",
		"ExecutionTypeParameter",
	)
	c.configuration = configuration
	return &c
}

func (c *ExecutionTypeParameter) GetValue() string {

	if c.configuration.GetExecutionType() == Proportional {
		return "Proportional"
	}
	return "Sequencial"
}

func (c *ExecutionTypeParameter) SetValue(value string) {
	value = strings.ToLower(value)
	if strings.HasPrefix(value, "p") {
		c.configuration.SetExecutionType(Proportional)
		return
	}
	c.configuration.SetExecutionType(Sequential)
}
