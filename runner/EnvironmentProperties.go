package runner

import (
	"os"

	utilities "github.com/pip-benchmark/pip-benchmark-go/utilities"
)

type EnvironmentProperties struct {
	*utilities.Properties
}

func NewEnvironmentProperties() *EnvironmentProperties {
	return &EnvironmentProperties{
		Properties: utilities.NewProperties(),
	}
}

func (c *EnvironmentProperties) getFilePath() string {
	return "./BenchmarkEnvironment.properties"
}

func (c *EnvironmentProperties) GetCpuBenchmark() float64 {
	return c.Properties.GetAsDouble("CpuBenchmark", 0)
}

func (c *EnvironmentProperties) SetCpuBenchmark(value float64) {
	c.Properties.SetAsDouble("CpuBenchmark", value)
}

func (c *EnvironmentProperties) GetDiskBenchmark() float64 {
	return c.Properties.GetAsDouble("DiskBenchmark", 0)
}

func (c *EnvironmentProperties) SetDiskBenchmark(value float64) {
	c.Properties.SetAsDouble("DiskBenchmark", value)
}

func (c *EnvironmentProperties) GetVideoBenchmark() float64 {
	return c.Properties.GetAsDouble("VideoBenchmark", 0)
}

func (c *EnvironmentProperties) SetVideoBenchmark(value float64) {
	c.Properties.SetAsDouble("VideoBenchmark", value)
}

func (c *EnvironmentProperties) load() {
	_, fileErr := os.Stat(c.getFilePath())
	if fileErr == nil { //os.IsExist(fileErr)
		c.Properties.LoadFromFile(c.getFilePath())
	}
}

func (c *EnvironmentProperties) Save() {
	c.Properties.SaveToFile(c.getFilePath())
}
