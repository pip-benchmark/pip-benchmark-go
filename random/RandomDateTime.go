package random

import (
	"time"
)

/*
Random generator for Date time values.

Example:

    value1 = RandomDateTime.NextDate(time.Date(2010,0,1));       // Possible result: 2008-01-03
    value2 = RandomDateTime.NextDateTime(time.Date(2017,0.1));   // Possible result: 2007-03-11 11:20:32
    value3 = RandomDateTime.UpdateDateTime(time.Date(2010,1,2)); // Possible result: 2010-02-05 11:33:23
*/
var RandomDateTime TRandomDateTime = TRandomDateTime{}

type TRandomDateTime struct {
}

// Generates a random Date in the range ['minYear', 'maxYear'].
// This method generate dates without time (or time set to 00:00:00)
// Parameters:
// - min   (optional) minimum range value
// - max   max range value
// Returns:     a random Date value.
func (c *TRandomDateTime) NextDate(min *time.Time, max *time.Time) time.Time {

	if max == nil {
		max = min
		tmp := time.Date(max.Year()-10, 1, 1, 0, 0, 0, 0, nil)
		min = &tmp
	}

	diff := max.UnixNano() - min.UnixNano()
	if diff <= 0 {
		return *min
	}

	tm := min.Unix() + RandomInteger.NextInteger(0, diff)
	date := time.Unix(tm, 0)

	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, nil)
}

// Generates a random Date and time in the range ['minYear', 'maxYear'].
// This method generate dates without time (or time set to 00:00:00)
// Parameters:
// - min   (optional) minimum range value
// - max   max range value
// Returns:     a random Date and time value.
func (c *TRandomDateTime) NextDateTime(min *time.Time, max *time.Time) time.Time {
	if max == nil {
		max = min
		tmp := time.Date(2000, 0, 1, 0, 0, 0, 0, nil)
		min = &tmp
	}

	diff := max.Unix() - min.Unix()
	if diff <= 0 {
		return *min
	}

	tm := min.Unix() + RandomInteger.NextInteger(0, diff)
	return time.Unix(tm, 0)
}

// Updates (drifts) a Date value within specified range defined
// - value     a Date value to drift.
// - range     (optional) a range in milliseconds. Default: 10 days
func (c *TRandomDateTime) UpdateDateTime(value time.Time, rng int64) time.Time {
	if rng == 0 {
		rng = 10 * 24 * 3600000
	}

	if rng < 0 {
		return value
	}

	tm := value.Unix() + RandomInteger.NextInteger(-rng, rng)
	return time.Unix(tm, 0)
}
