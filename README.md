# <img src="https://github.com/pip-devs/pip-selenium-ps/raw/master/artifacts/logo.png" alt="Pip.Services Logo" style="max-width:30%"> <br/> Portable Benchmarking Framework in Golang

This benchmarking framework ported cross multiple languages to obtain comparible performance metrics across different implementations.
In addition to performance benchmarking, it helps in other types of non-functional testing like load, reliability or concurrency.

* Measures performance in **transactions per second** or **TPS** 
* Supports **active** (by calling Execute method) or **passive** (by reporting via Context) measurement methods
* Supports **configuration parameters** to set connection strings or other settings for benchmarks
* Runs benchmarks **sequential** or in **proportional** by allocating % of calls to each benchmark
* Measures **peak** or **nominal** measurement at specified transaction rate
* Measures **system utilization** (RAM and CPU) during benchmarking process
* Measures overall **environment** performance (CPU, Video, Disk) for objective interpretation of results
* Capture and **errors** or **validation** results
* **Console runner** to execute benchmarks

## Usage

To run benchmark create your benchmark suit like this

```golang
package main

import (
	"errors"
	"time"

	bench "github.com/pip-benchmark/pip-benchmark-go/benchmark"
	benchconsole "github.com/pip-benchmark/pip-benchmark-go/console"
	rnd "github.com/pip-services3-go/pip-services3-commons-go/random"
)

//===============================
type SampleBenchmark struct {
	*bench.Benchmark
	greeting string
}

func NewSampleBenchmark() *SampleBenchmark {
    c := SampleBenchmark{}
    // Create base benchmark and set name with description
	c.Benchmark = bench.NewBenchmark("SampleBenchmark", "Measures performance of updating", "Type")
	c.Benchmark.IExecutable = &c
	c.greeting = "test"
	return &c
}

func (c *SampleBenchmark) SetUp() error {
    // Setup benckmark
    // This method must be always created
    // Get params from context
	c.greeting = c.Context.GetParameters()["Greeting"].GetAsString()
	return nil
}

func (c *SampleBenchmark) TearDown() error {
    // Teardown behckmark
    // This method must be always created
	return nil
}

func (c *SampleBenchmark) Execute() error {
    // Create testing function
    // This method must be always created
	// Randomly generate message or errors
	if rnd.RandomBoolean.Chance(1, 10) == true {
		c.Context.SendMessage(c.greeting)
	} else if rnd.RandomBoolean.Chance(1, 10) == true {
		c.Context.ReportError(errors.New("Something bad happend..."))
	}
	// Just wait and do nothing
	time.Sleep(10 * time.Millisecond)
	return nil
}

//===============================
type SampleBenchmarkSuite struct {
	*bench.BenchmarkSuite
}

func NewSampleBenchmarkSuite() *SampleBenchmarkSuite {
	c := SampleBenchmarkSuite{}
    c.BenchmarkSuite = bench.NewBenchmarkSuite("Samples", "Provides sample benchmarks")
    // Create parameters. They are will accessible in benchmarks from context
	c.CreateParameter("Greeting", "Greeting message", "Hello world!")
	c.AddBenchmark(NewSampleBenchmark().Benchmark)
	return &c
}
//===============================
func main() {
    // Create benchmark builder and setup it
	var benchmark = benchconsole.NewConsoleBenchmarkBuilder()
	benchmark.AddSuite(NewSampleBenchmarkSuite().BenchmarkSuite).
		ForDuration(5).
		ForceContinue(true).
		WithAllBenchmarks()
    runner := benchmark.Create()
    // You can change benchmarking params like this
	runner.Parameters().Set(map[string]string{"General.Benchmarking.MeasurementType": "Nominal"})
	runner.Parameters().Set(map[string]string{"General.Benchmarking.ExecutionType": "Sequential"})
    // Run bechmarking
	runner.Run(func(err error) {
		if err != nil {
			print(err.Error())
		}
	})
    // Print results
	print(runner.Report().Generate())
}

```



To measure environment (CPU, video, disk)
```bash
run ./app/main.go -e
```



## Acknowledgements

This module created and maintained by 
- **Sergey Seroukhov**
- **Levichev Dmitry**

