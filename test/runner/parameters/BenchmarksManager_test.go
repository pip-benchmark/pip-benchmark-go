package test_parameters

import (
	"testing"

	benchrunner "github.com/pip-benchmark/pip-benchmark-go/runner"
	"github.com/stretchr/testify/assert"
)

func TestBenchmarksManager(t *testing.T) {
	//t.Run("TestBenchmarksManager:LoadSuites", loadSuites)
	// t.Run("TestBenchmarksManager:AddSuiteFromClass", addSuiteFromClass)
	// t.Run("TestBenchmarksManager:SelectAll", selectAll)
	// t.Run("TestBenchmarksManager:SelectBenchmarkByName", selectBenchmarkByName)
}

func loadSuites(t *testing.T) {
	//test loadSuites
	runner := benchrunner.NewBenchmarkRunner()
	benchmarks := runner.Benchmarks()
	assert.Equal(t, 0, len(benchmarks.Suites()))
	benchmarks.AddSuitesFromModule("./obj/src/standardbenchmarks")
	assert.Equal(t, 2, len(benchmarks.Suites()))
}

func addSuiteFromClass(t *testing.T) {
	//test addSuiteFromClass
	runner := benchrunner.NewBenchmarkRunner()
	benchmarks := runner.Benchmarks()
	assert.Equal(t, 0, len(benchmarks.Suites()))
	benchmarks.AddSuiteFromClass("../../plugin/standardbenchmarks.so,UtilityBenchmarkSuite")
	assert.Equal(t, 1, len(benchmarks.Suites()))
}

func selectAll(t *testing.T) {
	//test selectAll
	runner := benchrunner.NewBenchmarkRunner()
	benchmarks := runner.Benchmarks()
	//benchmarks.AddSuitesFromModule("../../plugin/standardbenchmarks")
	benchmarks.AddSuiteFromClass("../../plugin/standardbenchmarks.so,UtilityBenchmarkSuite")
	benchmarks.AddSuiteFromClass("../../plugin/standardbenchmarks.so,StandardCpuBenchmark")
	benchmarks.AddSuiteFromClass("../../plugin/standardbenchmarks.so,StandardDiskBenchmark")
	benchmarks.AddSuiteFromClass("../../plugin/standardbenchmarks.so,StandardHardwareBenchmark")
	benchmarks.AddSuiteFromClass("../../plugin/standardbenchmarks.so,StandardVideoBenchmark")
	assert.Equal(t, 0, len(benchmarks.IsSelected()))
	runner.Benchmarks().SelectAll()
	assert.Equal(t, 5, len(benchmarks.IsSelected()))
	benchmarks.UnselectAll()
	assert.Equal(t, 0, len(benchmarks.IsSelected()))
}

func selectBenchmarkByName(t *testing.T) {
	//test selectBenchmarkByName
	runner := benchrunner.NewBenchmarkRunner()
	benchmarks := runner.Benchmarks()
	//benchmarks.AddSuitesFromModule("../../plugin/standardbenchmarks")
	benchmarks.AddSuiteFromClass("../../plugin/standardbenchmarks.so,UtilityBenchmarkSuite")
	benchmarks.AddSuiteFromClass("../../plugin/standardbenchmarks.so,StandardCpuBenchmark")
	benchmarks.AddSuiteFromClass("../../plugin/standardbenchmarks.so,StandardDiskBenchmark")
	benchmarks.AddSuiteFromClass("../../plugin/standardbenchmarks.so,StandardHardwareBenchmark")
	benchmarks.AddSuiteFromClass("../../plugin/standardbenchmarks.so,StandardVideoBenchmark")
	assert.Equal(t, 0, len(benchmarks.IsSelected()))
	benchmarks.SelectByName([]string{"Utility.Empty"})
	assert.Equal(t, 1, len(benchmarks.IsSelected()))
	benchmarks.UnselectByName([]string{"Utility.Empty"})
	assert.Equal(t, 0, len(benchmarks.IsSelected()))
}
