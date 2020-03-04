package benchmark

type Benchmark struct {
	IExecutable
	name        string
	description string
	Context     IExecutionContext
}

func NewBenchmark(name string, description string) *Benchmark {
	c := Benchmark{
		name:        name,
		description: description,
	}
	return &c
}

func (c *Benchmark) Name() string {
	return c.name
}

func (c *Benchmark) Description() string {
	return c.description
}

func (c *Benchmark) GetContext() IExecutionContext {
	return c.Context
}

func (c *Benchmark) SetContext(value IExecutionContext) {
	c.Context = value
}

func (c *Benchmark) SetUp() error {
	return nil
}

//func (c*Benchmark) abstract execute(callback: (err?: any) => void): void;

func (c *Benchmark) TearDown() error {
	return nil
}
