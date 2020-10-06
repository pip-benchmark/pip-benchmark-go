package benchmark

import (
	benchconv "github.com/pip-benchmark/pip-benchmark-go/utilities"
)

type Parameter struct {
	IParameter
	name         string
	description  string
	defaultValue string
	value        string
	paramType    string
}

func NewParameter(name string, description string, defaultValue string, paramType string) *Parameter {
	c := Parameter{
		name:         name,
		description:  description,
		defaultValue: defaultValue,
		value:        defaultValue,
		paramType:    paramType,
	}
	c.IParameter = &c
	return &c
}

func (c *Parameter) Type() string {
	return c.paramType
}

func (c *Parameter) Name() string {
	return c.name
}

func (c *Parameter) Description() string {
	return c.description
}

func (c *Parameter) DefaultValue() string {
	return c.defaultValue
}

func (c *Parameter) Value() string {
	return c.value
}

func (c *Parameter) SetValue(value string) {
	c.value = value
}

func (c *Parameter) GetAsString() string {
	return c.value
}

func (c *Parameter) GetAsNullableString() *string {
	if c.value == "" {
		return nil
	}
	buf := c.value
	return &buf
}

func (c *Parameter) GetAsStringWithDefault(defaultValue string) string {
	if c.value == "" {
		return defaultValue
	}
	return c.value
}

func (c *Parameter) SetAsString(value string) {
	c.value = value
}

func (c *Parameter) GetAsBoolean() bool {
	return benchconv.Converter.StringToBoolean(c.value, false)
}

func (c *Parameter) GetAsNullableBoolean() bool {
	return benchconv.Converter.StringToBoolean(c.value, false)
}

func (c *Parameter) GetAsBooleanWithDefault(defaultValue bool) bool {
	return benchconv.Converter.StringToBoolean(c.value, defaultValue)
}

func (c *Parameter) SetAsBoolean(value bool) {
	c.value = benchconv.Converter.BooleanToString(value)
}

func (c *Parameter) GetAsInteger() int {
	return benchconv.Converter.StringToInteger(c.value, 0)
}

func (c *Parameter) GetAsNullableInteger() int {
	return benchconv.Converter.StringToInteger(c.value, 0)
}

func (c *Parameter) GetAsIntegerWithDefault(defaultValue int) int {
	return benchconv.Converter.StringToInteger(c.value, defaultValue)
}

func (c *Parameter) SetAsInteger(value int) {
	c.value = benchconv.Converter.IntegerToString(value)
}

func (c *Parameter) GetAsLong() int32 {
	return benchconv.Converter.StringToLong(c.value, 0)
}

func (c *Parameter) GetAsNullableLong() int32 {
	return benchconv.Converter.StringToLong(c.value, 0)
}

func (c *Parameter) GetAsLongWithDefault(defaultValue int32) int32 {
	return benchconv.Converter.StringToLong(c.value, defaultValue)
}

func (c *Parameter) SetAsLong(value int32) {
	c.value = benchconv.Converter.LongToString(value)
}

func (c *Parameter) GetAsFloat() float32 {
	return benchconv.Converter.StringToFloat(c.value, 0)
}

func (c *Parameter) GetAsNullableFloat() float32 {
	return benchconv.Converter.StringToFloat(c.value, 0)
}

func (c *Parameter) GetAsFloatWithDefault(defaultValue float32) float32 {
	return benchconv.Converter.StringToFloat(c.value, defaultValue)
}

func (c *Parameter) SetAsFloat(value float32) {
	c.value = benchconv.Converter.FloatToString(value)
}

func (c *Parameter) GetAsDouble() float64 {
	return benchconv.Converter.StringToDouble(c.value, 0)
}

func (c *Parameter) GetAsNullableDouble() float64 {
	return benchconv.Converter.StringToDouble(c.value, 0)
}

func (c *Parameter) GetAsDoubleWithDefault(defaultValue float64) float64 {
	return benchconv.Converter.StringToDouble(c.value, defaultValue)
}

func (c *Parameter) SetAsDouble(value float64) {
	c.value = benchconv.Converter.DoubleToString(value)
}
