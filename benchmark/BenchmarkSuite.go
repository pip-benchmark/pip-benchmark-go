package benchmark

type BenchmarkSuite struct {
	name        string
	description string
	parameters  map[string]interface{}
	benchmarks  []*Benchmark
	context     IExecutionContext
}

func NewBenchmarkSuite(name string, description string) *BenchmarkSuite {
	c := BenchmarkSuite{
		name:        name,
		description: description,
	}
	c.benchmarks = make([]*Benchmark, 0)
	c.parameters = make(map[string]interface{}, 0)
	return &c
}

func (c *BenchmarkSuite) Name() string {
	return c.name
}

func (c *BenchmarkSuite) Description() string {
	return c.description
}

func (c *BenchmarkSuite) GetContext() IExecutionContext {
	return c.context
}

func (c *BenchmarkSuite) SetContext(value IExecutionContext) {
	c.context = value
}

func (c *BenchmarkSuite) Parameters() map[string]interface{} {
	return c.parameters
}

func (c *BenchmarkSuite) AddParameter(parameter Parameter) Parameter {
	c.parameters[parameter.name] = parameter
	return parameter
}

func (c *BenchmarkSuite) CreateParameter(name string, description string, defaultValue string) *Parameter {
	parameter := NewParameter(name, description, defaultValue, "")
	c.parameters[name] = parameter
	return parameter
}

func (c *BenchmarkSuite) Benchmarks() []*Benchmark {
	return c.benchmarks
}

func (c *BenchmarkSuite) AddBenchmark(benchmark *Benchmark) *Benchmark {
	c.benchmarks = append(c.benchmarks, benchmark)
	return benchmark
}

func (c *BenchmarkSuite) CreateBenchmark(name string, description string, executeCallback func() error) *Benchmark {
	benchmark := NewDelegatedBenchmark(name, description, executeCallback)
	c.benchmarks = append(c.benchmarks, benchmark.Benchmark)
	return benchmark.Benchmark
}

func (c *BenchmarkSuite) SetUp() error {
	return nil
}

func (c *BenchmarkSuite) TearDown() error {
	return nil
}
