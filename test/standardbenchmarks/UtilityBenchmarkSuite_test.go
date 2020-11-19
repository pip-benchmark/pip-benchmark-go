package test_standardbenchmarks

import (
	"sync"
	"testing"

	benchst "github.com/pip-benchmark/pip-benchmark-go/standardbenchmarks"
	"github.com/stretchr/testify/assert"
)

func TestUtilityBenchmarkSuite(t *testing.T) {
	var suite *benchst.UtilityBenchmarkSuite

	suite = benchst.NewUtilityBenchmarkSuite()
	suite.IPrepared.SetUp()

	defer suite.IPrepared.TearDown()

	//test emptyBenchmark
	assert.Equal(t, 2, len(suite.Benchmarks()))
	benchmark := suite.Benchmarks()[0]

	var wg sync.WaitGroup = sync.WaitGroup{}
	var err error
	wg.Add(1)

	go func() {
		defer wg.Done()

		for i := 0; i < 100; i++ {
			err = benchmark.Execute()
		}

	}()
	wg.Wait()
	assert.Nil(t, err)

	//  test randomBenchmark
	assert.Equal(t, 2, len(suite.Benchmarks()))
	benchmark = suite.Benchmarks()[1]
	wg.Add(1)

	go func() {
		defer wg.Done()

		for i := 0; i < 30; i++ {
			err = benchmark.Execute()
		}

	}()
	wg.Wait()
	assert.Nil(t, err)
}
