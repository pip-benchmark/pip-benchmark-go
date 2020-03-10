package test_standardbenchmarks

import (
	"sync"
	"testing"

	benchst "github.com/pip-benchmark/pip-benchmark-go/standardbenchmarks"
	"github.com/stretchr/testify/assert"
)

func TestStandardDiskBenchmark(t *testing.T) {
	var benchmark *benchst.StandardDiskBenchmark

	benchmark = benchst.NewStandardDiskBenchmark()
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
