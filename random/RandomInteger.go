package random

import (
	"math"
	"math/rand"
)

/*
Random generator for integer values.

Example:
    value1 = RandomInteger.NextInteger(5, 10);     // Possible result: 7
    value2 = RandomInteger.NextInteger(10);        // Possible result: 3
    value3 = RandomInteger.UpdateInteger(10, 3);   // Possible result: 9
*/
var RandomInteger TRandomInteger = TRandomInteger{}

type TRandomInteger struct {
}

// Generates a integer in the range ['min', 'max']. If 'max' is omitted, then the range will be set to [0, 'min'].
// Parameters:
// - min   minimum value of the integer that will be generated.
//              If 'max' is omitted, then 'max' is set to 'min' and 'min' is set to 0.
// - max   (optional) maximum value of the float that will be generated. Defaults to 'min' if omitted.
// Returns     generated random integer value.
func (c *TRandomInteger) NextInteger(min int64, max int64) int64 {
	if max == 0 {
		max = min
		min = 0
	}

	if max-min <= 0 {
		return min
	}

	return (int64)(math.Floor((float64)(min) + rand.Float64()*(float64)(max-min)))
}

// Updates (drifts) a integer value within specified range defined
// Parameters:
// - value     a integer value to drift.
// - range     (optional) a range. Default: 10% of the value
func (c *TRandomInteger) UpdateInteger(value int64, rng int64) int64 {

	if rng == 0 {
		rng = int64(math.Floor(0.1 * (float64)(value)))
	}
	minValue := value - rng
	maxValue := value + rng
	return c.NextInteger(minValue, maxValue)
}

// Generates a random sequence of integers starting from 0 like: [0,1,2,3...??]
// Parameters:
// - min   minimum value of the integer that will be generated.
//              If 'max' is omitted, then 'max' is set to 'min' and 'min' is set to 0.
// - max   (optional) maximum value of the float that will be generated. Defaults to 'min' if omitted.
// Returns:     generated array of integers.
func (c *TRandomInteger) Sequence(min int64, max int64) []int64 {

	if max == 0 {
		max = min
	}

	count := c.NextInteger(min, max)
	result := make([]int64, 0)
	for i := int64(0); i < count; i++ {
		result = append(result, i)
	}
	return result
}
