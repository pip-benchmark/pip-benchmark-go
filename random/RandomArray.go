package random

/*
Random generator for array objects.
Example:

    value1 := RandomArray.Pick([1, 2, 3, 4]); // Possible result: 3
*/
var RandomArray TRandomArray = TRandomArray{}

type TRandomArray struct {
}

// Picks a random element from specified array.
// - values    an array of any type
// Retruns:         a randomly picked item.

func (c *TRandomArray) Pick(values []interface{}) interface{} {
	if values == nil || len(values) == 0 {
		return nil
	}
	return values[RandomInteger.NextInteger(0, (int64)(len(values)))]
}
