package test_execution

import (
	"testing"
	"time"

	benchruner "github.com/pip-benchmark/pip-benchmark-go/runner"
	"github.com/stretchr/testify/assert"
)

func TestCpuLoadMeter(t *testing.T) {
	meter := benchruner.NewCpuLoadMeter()
	measure := meter.Measure()
	assert.True(t, measure == 0.0)

	select {
	case <-time.After(100 * time.Millisecond):
	}

	measure = meter.Measure()
	assert.True(t, measure > 0)
}
