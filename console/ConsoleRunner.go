package console

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	benchrunner "github.com/pip-benchmark/pip-benchmark-go/runner"
)

type ConsoleRunner struct {
	args   *CommandLineArgs
	runner *benchrunner.BenchmarkRunner
}

func (c *ConsoleRunner) start(args []string) {
	c.args = NewCommandLineArgs(args)
	c.runner = benchrunner.NewBenchmarkRunner()

	ConsoleEventPrinter.Attach(c.runner)

	c.executeBatchMode()
}

func (c *ConsoleRunner) Stop() {
	c.runner.Stop()
}

func (c *ConsoleRunner) executeBatchMode() {

	if c.args.ShowHelp {
		c.PrintHelp()
		return
	}

	// Load modules
	for _, module := range c.args.Modules {
		c.runner.Benchmarks().AddSuitesFromModule(module)
	}

	// Load test suites classes
	for _, class := range c.args.Classes {
		c.runner.Benchmarks().AddSuiteFromClass(class)
	}

	// Load configuration
	if c.args.ConfigurationFile != "" {
		c.runner.Parameters().LoadFromFile(c.args.ConfigurationFile)
	}

	// Set parameters
	if len(c.args.Parameters) != 0 {
		c.runner.Parameters().Set(c.args.Parameters)
	}

	// Select benchmarks
	if len(c.args.Benchmarks) == 0 {
		c.runner.Benchmarks().SelectAll()
	} else {
		c.runner.Benchmarks().SelectByName(c.args.Benchmarks)
	}

	if c.args.ShowParameters {
		c.printParameters()
		return
	}

	if c.args.ShowBenchmarks {
		c.printBenchmarks()
		return
	}
	var errGlobal error
	var wg sync.WaitGroup = sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Benchmark the environment
		if c.args.MeasureEnvironment {
			fmt.Println("Measuring Environment (wait up to 2 mins)...")
			msrErr := c.runner.Environment().Measure(true, true, false)
			output := fmt.Sprintf(
				"CPU: %s, Video: %s, Disk: %s",
				strconv.FormatFloat(c.runner.Environment().CpuMeasurement(), 'e', 2, 64),
				strconv.FormatFloat(c.runner.Environment().VideoMeasurement(), 'e', 2, 64),
				strconv.FormatFloat(c.runner.Environment().DiskMeasurement(), 'e', 2, 64))
			fmt.Println(output)
			if msrErr != nil {
				errGlobal = msrErr
			}

		}
	}()
	wg.Wait()

	wg.Add(1)

	go func() {
		defer wg.Done()
		// Configure benchmarking
		c.runner.Configuration().SetMeasurementType(c.args.MeasurementType)
		c.runner.Configuration().SetNominalRate(c.args.NominalRate)
		c.runner.Configuration().SetExecutionType(c.args.ExecutionType)
		c.runner.Configuration().SetDuration(c.args.Duration)

		// Perform benchmarking
		c.runner.Run(func(err error) {
			if len(c.runner.Results().All()) > 0 {
				fmt.Println(strconv.FormatFloat(c.runner.Results().All()[0].PerformanceMeasurement.AverageValue, 'e', 2, 64))
			}
			// Generate report and save to file
			if c.args.ReportFile != "" {
				c.runner.Report().SaveToFile(c.args.ReportFile)
			}

			// Show report in console
			if c.args.ShowReport {
				fmt.Println(c.runner.Report().Generate())
			}

			errGlobal = err
		})
	}()
	wg.Wait()

	if errGlobal != nil {
		fmt.Println(errGlobal)
	}

}

func (c *ConsoleRunner) PrintHelp() {
	fmt.Println("Pip.Benchmark Console Runner. (c) Conceptual Vision Consulting LLC 2017")
	fmt.Println()
	fmt.Println("Command Line Arguments:")
	fmt.Println("-a <module>    - Module with benchmarks to be loaded. You may include multiple modules")
	fmt.Println("-p <param>=<value> - Set parameter value. You may include multiple parameters")
	fmt.Println("-b <benchmark>   - Name of benchmark to be executed. You may include multiple benchmarks")
	fmt.Println("-c <config file> - File with parameters to be loaded")
	fmt.Println("-r <report file> - File to save benchmarking report")
	fmt.Println("-d <seconds>     - Benchmarking duration in seconds")
	fmt.Println("-h               - Display this help screen")
	fmt.Println("-B               - Show all available benchmarks")
	fmt.Println("-P               - Show all available parameters")
	fmt.Println("-R               - Show report")
	fmt.Println("-e               - Measure environment")
	fmt.Println("-x [prop|seq]    - Execution type: Proportional or Sequencial")
	fmt.Println("-m [peak|nominal] - Measurement type: Peak or Nominal")
	fmt.Println("-n <rate>        - Nominal rate in transactions per second")
}

func (c *ConsoleRunner) printBenchmarks() {
	fmt.Println("Pip.Benchmark Console Runner. (c) Conceptual Vision Consulting LLC 2017")
	fmt.Println()
	fmt.Println("Benchmarks:")

	suites := c.runner.Benchmarks().Suites()
	for _, suite := range suites {
		for _, benchmark := range suite.Benchmarks() {
			fmt.Println(benchmark.FullName() + " - " + benchmark.Description())
		}
	}
}

func (c *ConsoleRunner) printParameters() {
	fmt.Println("Pip.Benchmark Console Runner. (c) Conceptual Vision Consulting LLC 2017")
	fmt.Println()
	fmt.Println("Parameters:")

	parameters := c.runner.Parameters().UserDefined()
	for index := 0; index < len(parameters); index++ {
		parameter := parameters[index]
		defaultValue := parameter.DefaultValue()
		if defaultValue != "" {
			defaultValue = " (Default: " + defaultValue + ")"
		}
		fmt.Println(" " + parameter.Name() + " - " + parameter.Description() + defaultValue)
	}
}

func Run(args []string) {
	runner := &ConsoleRunner{}

	//         // Log uncaught exceptions
	//         process.on("uncaughtException", (ex) => {
	//             console.error(ex);
	//             console.error("Process is terminated");
	//             process.exit(1);
	//         });

	// Gracefully shutdown
	// process.on("exit", function () {
	//     runner.Stop();
	//     //fmt.Println("Goodbye!");
	// });
	if len(args) > 0 {
		runner.start(args)
	} else {
		runner.start(os.Args)
	}

}
