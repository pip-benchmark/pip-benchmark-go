package random

import (
	"math/rand"
	"time"

	ix "github.com/adam-lavrik/go-imath/ix"
)

/*
Random generator for boolean values.

Example:

    value1 := RandomBoolean.NextBoolean();    // Possible result: true
    value2 := RandomBoolean.Chance(1,3);      // Possible result: false
*/
var RandomBoolean TRandomBoolean = TRandomBoolean{}

type TRandomBoolean struct {
}

// Chance method are calculates "chance" out of "max chances".
// Example: 1 chance out of 3 chances (or 33.3%)
// - chance       a chance proportional to maxChances.
// - maxChances   a maximum number of chances
func (c *TRandomBoolean) Chance(chance int, maxChances int) bool {
	if chance < 0 {
		chance = 0
	}
	if maxChances < 0 {
		maxChances = 0
	}
	if chance == 0 && maxChances == 0 {
		return false
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	maxChances = ix.Max(maxChances, chance)
	start := (maxChances - chance) / 2
	end := start + chance
	hit := r.Float32() * (float32)(maxChances)
	return hit >= float32(start) && hit <= float32(end)
}

// Generates a random boolean value.
// Retruns: a random boolean.
func (c *TRandomBoolean) NextBoolean() bool {
	return rand.Intn(100) < 50
}
