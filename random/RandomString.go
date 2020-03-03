package random

import (
	"strings"
)

/*
Random generator for string values.
Example:

    value1 = RandomString.PickChar("ABC");     // Possible result: "C"
    value2 = RandomString.Pick(["A","B","C"]); // Possible result: "gBW"
*/
var RandomString TRandomString = NewTRandomString()

type TRandomString struct {
	digits     string
	symbols    string
	alphaLower string
	alphaUpper string
	alpha      string
	chars      string
}

func NewTRandomString() TRandomString {
	c := TRandomString{
		digits:     "01234956789",
		symbols:    "_,.:-/.[].{},#-!,$=%.+^.&*-() ",
		alphaLower: "abcdefghijklmnopqrstuvwxyz",
		alphaUpper: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}
	c.alpha = c.alphaUpper + c.alphaLower
	c.chars = c.alpha + c.digits + c.symbols
	return c
}

// Picks a random character from a string.
// Parameters:
//  - values    a string to pick a char from
// Returns     a randomly picked char.
func (c *TRandomString) PickChar(values string) string {
	if values == "" || len(values) == 0 {
		return string(0)
	}
	index := RandomInteger.NextInteger(0, int64(len(values)))
	return string(values[index])
}

// Picks a random string from an array of string.
// Parameters:
//     - values    strings to pick from.
// Returns         a randomly picked string.
func (c *TRandomString) Pick(values []string) string {
	if values == nil || len(values) == 0 {
		return ""
	}
	index := RandomInteger.NextInteger(0, int64(len(values)))
	return values[index]
}

// Distorts a string by randomly replacing characters in it.
// Parameters:
// - value    a string to distort.
// Returns        a distored string.
func (c *TRandomString) Distort(value string) string {

	value = strings.ToLower(value)
	//Capitalize the first letter of the string 'value'.
	if RandomBoolean.Chance(1, 5) {
		value = strings.ToUpper(string(value[0])) + value[1:]
	}
	//Add a symbol to the end of the string 'value'
	if RandomBoolean.Chance(1, 3) {
		value = value + string(c.PickChar(c.symbols))
	}
	return value
}

// Generates random alpha characted [A-Za-z]
// Returns a random characted.
func (c *TRandomString) NextAlphaChar() string {
	index := RandomInteger.NextInteger(0, int64(len(c.alpha)))
	return string(c.alpha[index])
}

// Generates a random string, consisting of upper and lower case letters (of the English alphabet),
// digits (0-9), and symbols ("_,.:-/.[].{},#-!,$=%.+^.&*-() ").
// Parameters:
// - minLength     (optional) minimum string length.
// - maxLength     maximum string length.
// Returns             a random string.
func (c *TRandomString) NextString(minLength int, maxLength int) string {
	result := ""

	length := RandomInteger.NextInteger(int64(minLength), int64(maxLength))
	for i := int64(0); i < length; i++ {
		index := RandomInteger.NextInteger(0, int64(len(c.alpha)))
		result += string(c.chars[index])
	}
	return result
}
