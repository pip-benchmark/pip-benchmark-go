package console

import (
	benchrunner "github.com/pip-benchmark/pip-benchmark-go/runner"
)

type ConsoleBenchmarkBuilder struct {
	*benchrunner.BenchmarkBuilder
}

func NewConsoleBenchmarkBuilder() *ConsoleBenchmarkBuilder {
	c := ConsoleBenchmarkBuilder{
		BenchmarkBuilder: benchrunner.NewBenchmarkBuilder(),
	}
	return &c
}

func (c *ConsoleBenchmarkBuilder) СonfigureWithArgs(args interface{}) *benchrunner.BenchmarkBuilder {
	var _args *CommandLineArgs

	if _, ok := args.(*CommandLineArgs); ok {
		_args = args.(*CommandLineArgs)
	} else {
		localArgs, ok := args.([]string)
		if !ok {
			panic("ConsoleBenchmarkBuilder: Error: Can't configure with args")
		}
		_args = NewCommandLineArgs(localArgs)
	}

	// Load modules
	for _, module := range _args.Modules {
		c.Runner.Benchmarks().AddSuitesFromModule(module)
	}

	// Load test suites classes
	for _, class := range _args.Classes {
		c.Runner.Benchmarks().AddSuiteFromClass(class)
	}

	// Load configuration
	if _args.ConfigurationFile != "" {
		c.Runner.Parameters().LoadFromFile(_args.ConfigurationFile)
	}

	// Set parameters
	if len(_args.Parameters) != 0 {
		c.Runner.Parameters().Set(_args.Parameters)
	}

	// Select benchmarks
	if len(_args.Benchmarks) == 0 {
		c.Runner.Benchmarks().SelectAll()
	} else {
		c.Runner.Benchmarks().SelectByName(_args.Benchmarks)
	}

	// Configure benchmarking
	c.Runner.Configuration().SetMeasurementType(_args.MeasurementType)
	c.Runner.Configuration().SetNominalRate(_args.NominalRate)
	c.Runner.Configuration().SetExecutionType(_args.ExecutionType)
	c.Runner.Configuration().SetDuration(_args.Duration)

	return c.BenchmarkBuilder
}

func (c *ConsoleBenchmarkBuilder) Create() *benchrunner.BenchmarkRunner {
	Runner := c.BenchmarkBuilder.Create()
	ConsoleEventPrinter.Attach(Runner)
	return Runner
}
