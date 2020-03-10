package test_standardbenchmarks

import (
	"sync"
	"testing"

	benchst "github.com/pip-benchmark/pip-benchmark-go/standardbenchmarks"
	"github.com/stretchr/testify/assert"
)

func TestStandardCpuBenchmark(t *testing.T) {
	var benchmark *benchst.StandardCpuBenchmark

	benchmark = benchst.NewStandardCpuBenchmark()
	benchmark.SetUp()

	defer benchmark.TearDown()
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
}
