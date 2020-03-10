package console

import (
	"strings"

	benchrunner "github.com/pip-benchmark/pip-benchmark-go/runner"
	util "github.com/pip-benchmark/pip-benchmark-go/utilities"
)

type CommandLineArgs struct {
	Modules            []string
	Classes            []string
	Benchmarks         []string
	Parameters         map[string]string
	ConfigurationFile  string
	ReportFile         string
	Duration           int64
	ShowHelp           bool
	ShowBenchmarks     bool
	ShowParameters     bool
	ShowReport         bool
	MeasureEnvironment bool
	MeasurementType    benchrunner.MeasurementType
	ExecutionType      benchrunner.ExecutionType
	NominalRate        float64
}

func NewCommandLineArgs(args []string) *CommandLineArgs {
	c := CommandLineArgs{}
	c.Modules = make([]string, 0)
	c.Classes = make([]string, 0)
	c.Benchmarks = make([]string, 0)
	c.Parameters = make(map[string]string, 0)
	c.ReportFile = "BenchmarkReport.txt"
	c.Duration = 60
	c.ShowHelp = false
	c.ShowBenchmarks = false
	c.ShowParameters = false
	c.ShowReport = false
	c.MeasureEnvironment = false
	c.MeasurementType = benchrunner.Peak
	c.ExecutionType = benchrunner.Proportional
	c.NominalRate = 1
	c.processArguments(args)
	return &c
}

func (c *CommandLineArgs) processArguments(args []string) {
	for index := 0; index < len(args); index++ {
		arg := args[index]
		moreArgs := index < len(args)-1

		if (arg == "-a" || arg == "-j" || arg == "--module") && moreArgs {
			index += 1
			module := args[index]
			c.Modules = append(c.Modules, module)
		} else if (arg == "-l" || arg == "--class") && moreArgs {
			index += 1
			class := args[index]
			c.Classes = append(c.Classes, class)
		} else if (arg == "-b" || arg == "--benchmark") && moreArgs {
			index += 1
			benchmark := args[index]
			c.Benchmarks = append(c.Benchmarks, benchmark)
		} else if (arg == "-p" || arg == "--param") && moreArgs {
			index += 1
			param := args[index]
			pos := strings.Index(param, "=")
			key := param
			if pos > 0 {
				key = param[:pos-1]
			}
			value := ""
			if pos > 0 {
				value = param[pos+1:]
			}
			c.Parameters[key] = value
		} else if (arg == "-c" || arg == "--config") && moreArgs {
			index += 1
			c.ConfigurationFile = args[index]
		} else if (arg == "-r" || arg == "--report") && moreArgs {
			index += 1
			c.ReportFile = args[index]
		} else if (arg == "-d" || arg == "--Duration") && moreArgs {
			index += 1
			c.Duration = int64(util.Converter.StringToLong(args[index], int32(60)))
		} else if (arg == "-m" || arg == "--measure") && moreArgs {
			index += 1
			measure := strings.ToLower(args[index])
			c.MeasurementType = benchrunner.Peak
			if strings.HasPrefix(measure, "nom") {
				c.MeasurementType = benchrunner.Nominal
			}
		} else if (arg == "-x" || arg == "--execute") && moreArgs {
			index += 1
			execution := strings.ToLower(args[index])
			c.ExecutionType = benchrunner.Proportional
			if strings.HasPrefix(execution, "seq") {
				c.ExecutionType = benchrunner.Sequential
			}

		} else if (arg == "-n" || arg == "--nominal") && moreArgs {
			index += 1
			c.NominalRate = util.Converter.StringToDouble(args[index], 1.0)
		} else if arg == "-h" || arg == "--help" {
			c.ShowHelp = true
		} else if arg == "-B" || arg == "--show-Benchmarks" {
			c.ShowBenchmarks = true
		} else if arg == "-P" || arg == "--show-params" {
			c.ShowParameters = true
		} else if arg == "-R" || arg == "--show-report" {
			c.ShowReport = true
		} else if arg == "-e" || arg == "--environment" {
			c.MeasureEnvironment = true
		}
	}
}
