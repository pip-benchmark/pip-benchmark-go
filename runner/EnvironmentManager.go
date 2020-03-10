package runner

import "sync"

type EnvironmentManager struct {
	*ExecutionManager
	duration         int64
	cpuMeasurement   float64
	videoMeasurement float64
	diskMeasurement  float64
}

func NewEnvironmentManager() *EnvironmentManager {
	c := EnvironmentManager{}
	c.duration = 5
	configuration := NewConfigurationManager()
	configuration.SetDuration(c.duration)
	results := NewResultsManager()
	c.ExecutionManager = NewExecutionManager(configuration, results)
	c.load()
	return &c
}

func (c *EnvironmentManager) SystemInfo() SystemInfo {
	return NewSystemInfo()
}

func (c *EnvironmentManager) CpuMeasurement() float64 {
	return c.cpuMeasurement
}

func (c *EnvironmentManager) VideoMeasurement() float64 {
	return c.videoMeasurement
}

func (c *EnvironmentManager) DiskMeasurement() float64 {
	return c.diskMeasurement
}

func (c *EnvironmentManager) Measure(cpu bool, disk bool, video bool) error {

	var errGlobal error
	var wg sync.WaitGroup = sync.WaitGroup{}

	wg.Add(1)
	go func() {
		if cpu {
			c.MeasureCpu(func(result float64, err error) {
				c.cpuMeasurement = result
				if err != nil {
					errGlobal = err
				}
				wg.Done()
			})
		}
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		if disk {
			c.MeasureDisk(func(result float64, err error) {
				c.diskMeasurement = result
				if err != nil {
					errGlobal = err
				}
				wg.Done()
			})
		}
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		if video {
			c.MeasureVideo(func(result float64, err error) {
				c.videoMeasurement = result
				if err != nil {
					errGlobal = err
				}
				wg.Done()
			})
		}
	}()

	wg.Wait()
	c.Stop()
	if errGlobal == nil {
		c.save()
	}
	return errGlobal
}

func (c *EnvironmentManager) load() {
	properties := NewEnvironmentProperties()
	properties.load()

	c.cpuMeasurement = properties.GetAsDouble("CpuMeasurement", 0)
	c.videoMeasurement = properties.GetAsDouble("VideoMeasurement", 0)
	c.diskMeasurement = properties.GetAsDouble("DiskMeasurement", 0)
}

func (c *EnvironmentManager) save() {
	properties := NewEnvironmentProperties()

	properties.SetAsDouble("CpuMeasurement", c.cpuMeasurement)
	properties.SetAsDouble("VideoMeasurement", c.videoMeasurement)
	properties.SetAsDouble("DiskMeasurement", c.diskMeasurement)

	properties.Save()
}

func (c *EnvironmentManager) MeasureCpu(callback func(result float64, err error)) {
	suite := NewStandardBenchmarkSuite()
	instance := NewBenchmarkSuiteInstance(suite.BenchmarkSuite)

	instance.UnselectAll()
	instance.SelectByName(suite.cpuBenchmark.Name())

	c.Run(instance.IsSelected(), func(err error) {
		result := 0.0
		if len(c.Results.All()) > 0 {
			result = c.Results.All()[0].PerformanceMeasurement.AverageValue
		}
		callback(result, err)
	})
}

func (c *EnvironmentManager) MeasureDisk(callback func(result float64, err error)) {
	suite := NewStandardBenchmarkSuite()
	instance := NewBenchmarkSuiteInstance(suite.BenchmarkSuite)

	instance.UnselectAll()
	instance.SelectByName(suite.diskBenchmark.Name())

	c.Run(instance.IsSelected(), func(err error) {
		result := 0.0
		if len(c.Results.All()) > 0 {
			result = c.Results.All()[0].PerformanceMeasurement.AverageValue
		}
		callback(result, err)
	})
}

func (c *EnvironmentManager) MeasureVideo(callback func(result float64, err error)) {
	suite := NewStandardBenchmarkSuite()
	instance := NewBenchmarkSuiteInstance(suite.BenchmarkSuite)

	instance.UnselectAll()
	instance.SelectByName(suite.videoBenchmark.Name())

	c.Run(instance.IsSelected(), func(err error) {
		result := 0.0
		if len(c.Results.All()) > 0 {
			result = c.Results.All()[0].PerformanceMeasurement.AverageValue
		}
		callback(result, err)
	})
}
