package test_utilities

import (
	"testing"

	benchutil "github.com/pip-benchmark/pip-benchmark-go/utilities"
	"github.com/stretchr/testify/assert"
)

func TestConverter(t *testing.T) {
	//test longToString
	//assert.Equal(t, "", benchutil.Converter.LongToString(""))
	assert.Equal(t, "123", benchutil.Converter.LongToString(123))

	//test stringToLong
	//assert.Equal(t, 0, benchutil.Converter.StringToLong("", 0))
	assert.Equal(t, int32(0), benchutil.Converter.StringToLong("ABC", 0))
	assert.Equal(t, int32(123), benchutil.Converter.StringToLong("123", 0))

	//test doubleToString
	//assert.Equal(t, "", benchutil.Converter.DoubleToString(""))
	assert.Equal(t, "123.456", benchutil.Converter.DoubleToString(123.456))

	//test stringToDouble
	//assert.Equal(t, 0, benchutil.Converter.StringToDouble("", 0))
	assert.Equal(t, 0.0, benchutil.Converter.StringToDouble("ABC", 0))
	assert.Equal(t, 123.456, benchutil.Converter.StringToDouble("123.456", 0))

	//test booleanToString
	//assert.Equal(t, "false", benchutil.Converter.BooleanToString(""))
	assert.Equal(t, "true", benchutil.Converter.BooleanToString(true))

	//test stringToBoolean
	//assert.Equal(t, false, benchutil.Converter.StringToBoolean("", false))
	assert.Equal(t, true, benchutil.Converter.StringToBoolean("True", false))
	assert.Equal(t, true, benchutil.Converter.StringToBoolean("1", false))
	assert.Equal(t, true, benchutil.Converter.StringToBoolean("T", false))
}
