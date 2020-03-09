package runner

import (
	"strings"

	bench "github.com/pip-benchmark/pip-benchmark-go/benchmark"
	benchutil "github.com/pip-benchmark/pip-benchmark-go/utilities"
)

type ParametersManager struct {
	configuration *ConfigurationManager
	parameters    []*bench.Parameter
}

func NewParametersManager(configuration *ConfigurationManager) *ParametersManager {
	c := ParametersManager{}
	c.configuration = configuration
	c.parameters = make([]*bench.Parameter, 0)
	c.parameters = append(c.parameters, NewMeasurementTypeParameter(configuration).Parameter)
	c.parameters = append(c.parameters, NewNominalRateParameter(configuration).Parameter)
	c.parameters = append(c.parameters, NewExecutionTypeParameter(configuration).Parameter)
	c.parameters = append(c.parameters, NewDurationParameter(configuration).Parameter)
	return &c
}

func (c *ParametersManager) UserDefined() []*bench.Parameter {
	var parameters []*bench.Parameter

	for _, parameter := range c.parameters {
		if !strings.HasSuffix(parameter.Name(), ".Selected") &&
			!strings.HasSuffix(parameter.Name(), ".Proportion") &&
			!strings.HasPrefix(parameter.Name(), "General.") {
			parameters = append(parameters, parameter)
		}
	}

	return parameters
}

func (c *ParametersManager) All() []*bench.Parameter {
	return c.parameters
}

func (c *ParametersManager) LoadFromFile(path string) {
	properties := benchutil.Properties{}
	properties.LoadFromFile(path)

	for _, parameter := range c.parameters {
		prop := properties.GetAsString(parameter.Name(), "")
		if prop != "" {
			parameter.SetValue(prop)
		}
	}

	c.configuration.NotifyChanged()
}

func (c *ParametersManager) SaveToFile(path string) {
	properties := benchutil.Properties{}
	for _, parameter := range c.parameters {
		properties.SetAsString(parameter.Name(), parameter.Value())
	}
	properties.SaveToFile(path)
}

func (c *ParametersManager) AddSuite(suite *BenchmarkSuiteInstance) {
	for _, benchmark := range suite.Benchmarks() {
		benchmarkSelectedParameter := NewBenchmarkSelectedParameter(benchmark)
		c.parameters = append(c.parameters, benchmarkSelectedParameter.Parameter)

		benchmarkProportionParameter := NewBenchmarkProportionParameter(benchmark)
		c.parameters = append(c.parameters, benchmarkProportionParameter.Parameter)
	}

	for _, parameter := range suite.Parameters() {
		suiteParameter := NewBenchmarkSuiteParameter(suite, &parameter)
		c.parameters = append(c.parameters, suiteParameter.Parameter)
	}

	c.configuration.NotifyChanged()
}

func (c *ParametersManager) RemoveSuite(suite *BenchmarkSuiteInstance) {
	parameterNamePrefix := suite.Name() + "."

	for index, param := range c.parameters {
		if strings.HasPrefix(param.Name(), parameterNamePrefix) {
			if index == len(c.parameters) {
				c.parameters = c.parameters[:index-1]
			} else {
				c.parameters = append(c.parameters[:index], c.parameters[index+1:]...)
			}
		}
	}
	c.configuration.NotifyChanged()
}

func (c *ParametersManager) SetToDefault() {
	for _, parameter := range c.parameters {
		if parameter.Type() == "BenchmarkSuiteParameter" {
			parameter.SetValue(parameter.DefaultValue())
		}
	}
	c.configuration.NotifyChanged()
}

func (c *ParametersManager) Set(parameters map[string]string) {
	for _, parameter := range c.parameters {
		if param, ok := parameters[parameter.Name()]; ok {
			parameter.SetValue(param)
		}
	}
	c.configuration.NotifyChanged()
}
