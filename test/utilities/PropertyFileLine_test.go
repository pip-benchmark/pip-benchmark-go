package test_utilities

import (
	"testing"

	benchutil "github.com/pip-benchmark/pip-benchmark-go/utilities"
	"github.com/stretchr/testify/assert"
)

func TestPropertyFileLine(t *testing.T) {

	//test compose
	line := benchutil.NewPropertyFileLine("Key", "Value", "Comment")
	assert.Equal(t, "Key", line.Key())
	assert.Equal(t, "Value", line.GetValue())
	assert.Equal(t, "Comment", line.GetComment())
	assert.Equal(t, "Key=Value ;Comment", line.Line())

	//test parse key
	line = benchutil.NewPropertyFileLine("Key", "", "")
	assert.Equal(t, "Key", line.Key())
	assert.Equal(t, "", line.GetValue())
	assert.Equal(t, "", line.GetComment())

	//test parse key/value
	line = benchutil.NewPropertyFileLine("Key=Value", "", "")
	assert.Equal(t, "Key", line.Key())
	assert.Equal(t, "Value", line.GetValue())
	assert.Equal(t, "", line.GetComment())

	line = benchutil.NewPropertyFileLine("Key='Value'", "", "")
	assert.Equal(t, "Key", line.Key())
	assert.Equal(t, "Value", line.GetValue())
	assert.Equal(t, "", line.GetComment())

	line = benchutil.NewPropertyFileLine(`Key="Value"`, "", "")
	assert.Equal(t, "Key", line.Key())
	assert.Equal(t, "Value", line.GetValue())
	assert.Equal(t, "", line.GetComment())

	//test parse full line
	line = benchutil.NewPropertyFileLine("Key=Value;Comment", "", "")
	assert.Equal(t, "Key", line.Key())
	assert.Equal(t, "Value", line.GetValue())
	assert.Equal(t, "Comment", line.GetComment())

}
