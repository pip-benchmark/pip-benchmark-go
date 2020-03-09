package console

import (
	"fmt"
	"strconv"
	"time"

	benchrunner "github.com/pip-benchmark/pip-benchmark-go/runner"
)

var ConsoleEventPrinter TConsoleEventPrinter = TConsoleEventPrinter{}

type TConsoleEventPrinter struct {
}

func (c *TConsoleEventPrinter) Attach(runner *benchrunner.BenchmarkRunner) {
	upd := benchrunner.ExecutionCallback(ConsoleEventPrinter.OnStateUpdated)
	runner.Execution().AddUpdatedListener(&upd)
	errLis := benchrunner.ErrorCallback(ConsoleEventPrinter.OnErrorReported)
	runner.Results().AddErrorListener(&errLis)
	msgLis := benchrunner.MessageCallback(ConsoleEventPrinter.OnMessageSent)
	runner.Results().AddMessageListener(&msgLis)
	resUpd := benchrunner.ResultCallback(ConsoleEventPrinter.OnResultUpdated)
	runner.Results().AddUpdatedListener(&resUpd)
}

func (c *TConsoleEventPrinter) OnStateUpdated(state benchrunner.ExecutionState) {
	if state == benchrunner.Running {
		fmt.Println("Measuring....")
	} else if state == benchrunner.Completed {
		fmt.Println("Completed measuring.")
	}
}

func (c *TConsoleEventPrinter) OnResultUpdated(result *benchrunner.BenchmarkResult) {
	if result != nil {
		output := fmt.Sprintf("%s Performance: %s %s>%s>%s CPU Load: %s %s>%s>%s Errors: %d",
			time.Now().Format(time.RFC3339),
			strconv.FormatFloat(result.PerformanceMeasurement.CurrentValue, 'e', 2, 64),
			strconv.FormatFloat(result.PerformanceMeasurement.MinValue, 'e', 2, 64),
			strconv.FormatFloat(result.PerformanceMeasurement.AverageValue, 'e', 2, 64),
			strconv.FormatFloat(result.PerformanceMeasurement.MaxValue, 'e', 2, 64),
			strconv.FormatFloat(result.CpuLoadMeasurement.CurrentValue, 'e', 2, 64),
			strconv.FormatFloat(result.CpuLoadMeasurement.MinValue, 'e', 2, 64),
			strconv.FormatFloat(result.CpuLoadMeasurement.AverageValue, 'e', 2, 64),
			strconv.FormatFloat(result.CpuLoadMeasurement.MaxValue, 'e', 2, 64),
			len(result.Errors),
		)
		fmt.Println(output)
	}
}

func (c *TConsoleEventPrinter) OnMessageSent(message string) {
	fmt.Println(message)
}

func (c *TConsoleEventPrinter) OnErrorReported(message error) {
	fmt.Println(message)
}
