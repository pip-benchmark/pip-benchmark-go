package utilities

import (
	"strconv"
	"strings"
)

var Converter TConverter = TConverter{}

type TConverter struct {
}

func (c *TConverter) StringToInteger(value string, defaultValue int) int {
	return int(c.StringToLong(value, int32(defaultValue)))
}

func (c *TConverter) IntegerToString(value int) string {
	return c.LongToString(int32(value))
}

func (c *TConverter) StringToLong(value string, defaultValue int32) int32 {
	if value == "" {
		return defaultValue
	}
	res, convErr := strconv.ParseInt(value, 10, 32)
	if convErr != nil {
		return defaultValue
	}
	return int32(res)
}

func (c *TConverter) LongToString(value int32) string {
	return strconv.FormatInt(int64(value), 10)
}

func (c *TConverter) StringToFloat(value string, defaultValue float32) float32 {
	return float32(c.StringToDouble(value, float64(defaultValue)))
}

func (c *TConverter) FloatToString(value float32) string {
	return c.DoubleToString(float64(value))
}

func (c *TConverter) StringToDouble(value string, defaultValue float64) float64 {
	if value == "" {
		return defaultValue
	}
	result, convErr := strconv.ParseFloat(value, 64)
	if convErr != nil {
		return defaultValue
	}
	return result
}

func (c *TConverter) DoubleToString(value float64) string {
	return strconv.FormatFloat(value, 'e', 5, 64)
}

func (c *TConverter) StringToBoolean(value string, defaultValue bool) bool {
	// Process nulls or empty strings
	if value == "" {
		return defaultValue
	}

	// Process single characters
	if len(value) == 1 {
		if value[0] == '1' || value[0] == 'T' || value[0] == 'Y' ||
			value[0] == 't' || value[0] == 'y' {
			return true
		}
		if value[0] == '0' || value[0] == 'F' || value[0] == 'N' ||
			value[0] == 'f' || value[0] == 'n' {
			return false
		}
	}

	// Process strings
	value = strings.ToUpper(value)
	if value == "TRUE" || value == "YES" {
		return true
	}
	if value == "FALSE" || value == "NO" {
		return false
	}

	return defaultValue
}

func (c *TConverter) BooleanToString(value bool) string {
	if value {
		return "true"
	}
	return "false"
}
