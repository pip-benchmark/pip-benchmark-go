package benchmark

type Benchmark struct {
	IExecutionContext
	name        string
	description string
	context     IExecutionContext
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
	return c.context
}

func (c *Benchmark) SetContext(value IExecutionContext) {
	c.context = value
}

func (c *Benchmark) SetUp() error {
	return nil
}

//func (c*Benchmark) abstract execute(callback: (err?: any) => void): void;

func (c *Benchmark) TearDown() error {
	return nil
}
