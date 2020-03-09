package environment

import (
	benchconf "github.com/pip-benchmark/pip-benchmark-go/runner/config"
	benchexec "github.com/pip-benchmark/pip-benchmark-go/runner/execution"
	benchres "github.com/pip-benchmark/pip-benchmark-go/runner/results"
)

// import { ConfigurationManager } from '../config/ConfigurationManager';
// import { ResultsManager } from '../results/ResultsManager';
// import { ExecutionManager } from '../execution/ExecutionManager';
// import { EnvironmentProperties } from './EnvironmentProperties';
// import { BenchmarkSuiteInstance } from '../benchmarks/BenchmarkSuiteInstance';
// import { StandardBenchmarkSuite } from './StandardBenchmarkSuite';
// import { SystemInfo } from './SystemInfo';

type EnvironmentManager struct {
	*benchexec.ExecutionManager
	duration         int64
	cpuMeasurement   float64
	videoMeasurement float64
	diskMeasurement  float64
}

func NewEnvironmentManager() *EnvironmentManager {
	c := EnvironmentManager{}
	c.duration = 5
	configuration := benchconf.NewConfigurationManager()
	configuration.SetDuration(c.duration)
	results := benchres.NewResultsManager()
	c.ExecutionManager = benchexec.NewExecutionManager(configuration, results)
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

// func (c*EnvironmentManager) Measure(cpu: boolean, disk: boolean, video: boolean, callback?: (err: any) => void) {
//     async.series([
//         (callback) => {
//             if (cpu) {
//                 c.measureCpu((err, result) => {
//                     if (err == null)
//                         c.cpuMeasurement = result;
//                     callback();
//                 })
//             } else callback();
//         },
//         (callback) => {
//             if (disk) {
//                 c.measureDisk((err, result) => {
//                     if (err == null)
//                         c.diskMeasurement = result;
//                     callback();
//                 })
//             } else callback();
//         },
//         (callback) => {
//             if (video) {
//                 c.measureVideo((err, result) => {
//                     if (err == null)
//                         c.videoMeasurement = result;
//                     callback();
//                 })
//             } else callback();
//         }
//     ], (err) => {
//         c.stop();

//         if (err == null)
//             c.save();

//         if (callback) callback(err);
//     });
// }

func (c *EnvironmentManager) load() {
	properties := EnvironmentProperties{}
	properties.load()

	c.cpuMeasurement = properties.GetAsDouble("CpuMeasurement", 0)
	c.videoMeasurement = properties.GetAsDouble("VideoMeasurement", 0)
	c.diskMeasurement = properties.GetAsDouble("DiskMeasurement", 0)
}

func (c *EnvironmentManager) save() {
	properties := EnvironmentProperties{}

	properties.SetAsDouble("CpuMeasurement", c.cpuMeasurement)
	properties.SetAsDouble("VideoMeasurement", c.videoMeasurement)
	properties.SetAsDouble("DiskMeasurement", c.diskMeasurement)

	properties.Save()
}

// func (c *EnvironmentManager) measureCpu(callback: (err: any, result: number) => void) {
//     let suite = new StandardBenchmarkSuite();
//     let instance = new BenchmarkSuiteInstance(suite);

//     instance.unselectAll();
//     instance.selectByName(suite.cpuBenchmark.name);

//     c.run(instance.isSelected, (err) => {
//         let result = c._results.all.length > 0
//             ? c._results.all[0].performanceMeasurement.averageValue : 0;
//         callback(err, result);
//     });
// }

// func (c *EnvironmentManager) measureDisk(callback: (err: any, result: number) => void) {
//     let suite = new StandardBenchmarkSuite();
//     let instance = new BenchmarkSuiteInstance(suite);

//     instance.unselectAll();
//     instance.selectByName(suite.diskBenchmark.name);

//     c.run(instance.isSelected, (err) => {
//         let result = c._results.all.length > 0
//             ? c._results.all[0].performanceMeasurement.averageValue : 0;
//         callback(err, result);
//     });
// }

// func (c *EnvironmentManager) measureVideo(callback: (err: any, result: number) => void) {
//     let suite = new StandardBenchmarkSuite();
//     let instance = new BenchmarkSuiteInstance(suite);

//     instance.unselectAll();
//     instance.selectByName(suite.videoBenchmark.name);

//     c.run(instance.isSelected, (err) => {
//         let result = c._results.all.length > 0
//             ? c._results.all[0].performanceMeasurement.averageValue : 0;
//         callback(err, result);
//     });
// }
