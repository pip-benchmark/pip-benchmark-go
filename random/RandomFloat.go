package random

import (
	"math/rand"
)

/*
Random generator for float values.

Example:

     value1 = RandomFloat.NextFloat(5, 10);     // Possible result: 7.3
     value2 = RandomFloat.NextFloat(0, 10);     // Possible result: 3.7
     value3 = RandomFloat.UpdateFloat(10, 3);   // Possible result: 9.2
*/
var RandomFloat TRandomFloat = TRandomFloat{}

type TRandomFloat struct{}

// Generates a float in the range ['min', 'max']. If 'max' is omitted, then the range will be set to [0, 'min'].
// Parameters:
//  - min   minimum value of the float that will be generated.
//               If 'max' is omitted, then 'max' is set to 'min' and 'min' is set to 0.
//  - max   (optional) maximum value of the float that will be generated. Defaults to 'min' if omitted.
// Returns:     generated random float value.
func (c *TRandomFloat) NextFloat(min float32, max float32) float32 {
	if max == 0 {
		max = min
		min = 0
	}

	if max-min <= 0 {
		return min
	}

	return min + rand.Float32()*(max-min)
}

// Updates (drifts) a float value within specified range defined

//  - value     a float value to drift.
//  - range     (optional) a range. Default: 10% of the value
func (c *TRandomFloat) UpdateFloat(value float32, rng float32) float32 {
	if rng == 0 {
		rng = 0.1 * value
	}
	minValue := value - rng
	maxValue := value + rng
	return RandomFloat.NextFloat(minValue, maxValue)
}
