package random

import (
	"math/rand"
)

/*
Random generator for double values.

Example:

    value1 = RandomDouble.NextDouble(5, 10);     // Possible result: 7.3
    value2 = RandomDouble.NextDouble(0, 10);        // Possible result: 3.7
    value3 = RandomDouble.UpdateDouble(10, 3);   // Possible result: 9.2
*/
var RandomDouble TRandomDouble = TRandomDouble{}

type TRandomDouble struct {
}

// NextDouble are generates a random double value in the range ['minYear', 'maxYear'].
// Paramaters:
// - min   (optional) minimum range value
// - max   max range value
// Returns:     a random double value.
func (c *TRandomDouble) NextDouble(min float64, max float64) float64 {
	if max == 0 {
		max = min
		min = 0
	}
	if max-min <= 0 {
		return min
	}
	return min + rand.Float64()*(max-min)
}

// UpdateDouble method are updates (drifts) a double value within specified range defined
// Parameters:
// - value     a double value to drift.
// - range     (optional) a range. Default: 10% of the value
// Returns:     a random double value.
func (c *TRandomDouble) UpdateDouble(value float64, rng float64) float64 {
	if rng == 0 {
		rng = 0.1 * value
	}
	minValue := value - rng
	maxValue := value + rng
	return RandomDouble.NextDouble(minValue, maxValue)
}
