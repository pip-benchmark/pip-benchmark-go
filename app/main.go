package main

import (
	"errors"
	"time"

	bench "github.com/pip-benchmark/pip-benchmark-go/benchmark"
	benchconsole "github.com/pip-benchmark/pip-benchmark-go/console"
	rnd "github.com/pip-services3-go/pip-services3-commons-go/random"
)

type SampleBenchmark struct {
	*bench.Benchmark
	greeting string
}

func NewSampleBenchmark() *SampleBenchmark {
	c := SampleBenchmark{}
	c.Benchmark = bench.NewBenchmark("SampleBenchmark", "Measures performance of updating", "Type")
	c.Benchmark.IExecutable = &c
	c.greeting = "test"
	return &c
}

func (c *SampleBenchmark) SetUp() error {
	c.greeting = c.Context.GetParameters()["Greeting"].GetAsString()
	return nil
}

func (c *SampleBenchmark) TearDown() error {
	return nil
}

func (c *SampleBenchmark) Execute() error {
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

type SampleBenchmarkSuite struct {
	*bench.BenchmarkSuite
}

func NewSampleBenchmarkSuite() *SampleBenchmarkSuite {
	c := SampleBenchmarkSuite{}
	c.BenchmarkSuite = bench.NewBenchmarkSuite("Samples", "Provides sample benchmarks")
	c.CreateParameter("Greeting", "Greeting message", "Hello world!")
	c.AddBenchmark(NewSampleBenchmark().Benchmark)
	return &c
}

func main() {
	//benchconsole.Run(os.Args)
	//benchconsole.Run([]string{"-e"})

	var benchmark = benchconsole.NewConsoleBenchmarkBuilder()
	benchmark.AddSuite(NewSampleBenchmarkSuite().BenchmarkSuite).
		ForDuration(5).
		ForceContinue(true).
		WithAllBenchmarks()
	runner := benchmark.Create()
	runner.Parameters().Set(map[string]string{"General.Benchmarking.MeasurementType": "Nominal"})
	runner.Parameters().Set(map[string]string{"General.Benchmarking.ExecutionType": "Sequential"})

	runner.Run(func(err error) {
		if err != nil {
			print(err.Error())
		}
	})

	print(runner.Report().Generate())
}
