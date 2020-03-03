package random

import (
	"strconv"
	"strings"

	ix "github.com/adam-lavrik/go-imath/ix"
)

/*
Random generator for various text values like names, addresses or phone numbers.

Example:

     value1 = c.Name();     // Possible result: "Segio"
     value2 = c.Verb();      // Possible result: "Run"
     value3 = c.Text(50);    // Possible result: "Run jorge. Red high scream?"
*/
var RandomText TRandomText = NewTRandomText()

type TRandomText struct {
	namePrefixes []string
	nameSuffixes []string
	firstNames   []string
	lastNames    []string
	colors       []string
	stuffs       []string
	adjectives   []string
	verbs        []string
	allWords     []string
}

func NewTRandomText() TRandomText {
	c := TRandomText{}
	c.namePrefixes = []string{"Dr.", "Mr.", "Mrs"}
	c.nameSuffixes = []string{"Jr.", "Sr.", "II", "III"}
	c.firstNames = []string{
		"John", "Bill", "Andrew", "Nick", "Pamela", "Bela", "Sergio", "George", "Hurry", "Cecilia", "Vesta", "Terry", "Patrick",
	}
	c.lastNames = []string{
		"Doe", "Smith", "Johns", "Gates", "Carmack", "Zontak", "Clinton", "Adams", "First", "Lopez", "Due", "White", "Black",
	}
	c.colors = []string{
		"Black", "White", "Red", "Blue", "Green", "Yellow", "Purple", "Grey", "Magenta", "Cian",
	}
	c.stuffs = []string{
		"Game", "Ball", "Home", "Board", "Car", "Plane", "Hotel", "Wine", "Pants", "Boots", "Table", "Chair",
	}
	c.adjectives = []string{
		"Large", "Small", "High", "Low", "Certain", "Fuzzy", "Modern", "Faster", "Slower",
	}
	c.verbs = []string{
		"Run", "Stay", "Breeze", "Fly", "Lay", "Write", "Draw", "Scream",
	}

	c.allWords = append(c.allWords, c.firstNames...)
	c.allWords = append(c.allWords, c.lastNames...)
	c.allWords = append(c.allWords, c.colors...)
	c.allWords = append(c.allWords, c.stuffs...)
	c.allWords = append(c.allWords, c.adjectives...)
	c.allWords = append(c.allWords, c.verbs...)
	return c
}

// Generates a random color name.
// The result value is capitalized.
// Returns: a random color name.
func (c *TRandomText) Color() string {
	return RandomString.Pick(c.colors)
}

// Generates a random noun.
// The result value is capitalized.
// Returns: a random noun.
func (c *TRandomText) Noun() string {
	return RandomString.Pick(c.stuffs)
}

// Generates a random adjective.
// The result value is capitalized.
// Returns: a random adjective.
func (c *TRandomText) Adjective() string {
	return RandomString.Pick(c.adjectives)
}

// Generates a random verb.
// The result value is capitalized.
// Returns: a random verb.
func (c *TRandomText) Verb() string {
	return RandomString.Pick(c.verbs)
}

// Generates a random phrase which consists of few words separated by spaces.
// The first word is capitalized, others are not.
//   - minLength     (optional) minimum string length.
//   - maxLength     maximum string length.
// Returns: a random phrase.
func (c *TRandomText) Phrase(minLength int, maxLength int) string {
	maxLength = ix.Max(minLength, maxLength)
	size := RandomInteger.NextInteger(int64(minLength), int64(maxLength))
	if size <= 0 {
		return ""
	}

	result := ""
	result += RandomString.Pick(c.allWords)
	for int64(len(result)) < size {
		result += " " + strings.ToLower(RandomString.Pick(c.allWords))
	}
	return result
}

// Generates a random person"s name which has the following structure
// <optional prefix> <first name> <second name> <optional suffix>
// Returns: a random name.
func (c *TRandomText) FullName() string {
	result := ""

	if RandomBoolean.Chance(3, 5) {
		result += RandomString.Pick(c.namePrefixes) + " "
	}
	result += RandomString.Pick(c.firstNames) + " " + RandomString.Pick(c.lastNames)
	if RandomBoolean.Chance(5, 10) {
		result += " " + RandomString.Pick(c.nameSuffixes)
	}
	return result
}

// Generates a random word from available first names, last names, colors, stuffs, adjectives, or verbs.
// Returns: a random word.
func (c *TRandomText) Word() string {
	return RandomString.Pick(c.allWords)
}

// Generates a random text that consists of random number of random words separated by spaces.
// - min   (optional) a minimum number of words.
// - max   a maximum number of words.
// Returns:     a random text.
func (c *TRandomText) Words(min int, max int) string {
	result := ""

	count := (int)(RandomInteger.NextInteger(int64(min), int64(max)))
	for i := 0; i < count; i++ {
		result += RandomString.Pick(c.allWords)
	}
	return result
}

// Generates a random phone number.
// The phone number has the format: (XXX) XXX-YYYY
// Returns: a random phone number.
func (c *TRandomText) Phone() string {
	result := ""

	result += "(" +
		strconv.FormatInt(RandomInteger.NextInteger(111, 999), 10) +
		") " +
		strconv.FormatInt(RandomInteger.NextInteger(111, 999), 10) +
		"-" +
		strconv.FormatInt(RandomInteger.NextInteger(0, 9999), 10)
	return result
}

// Generates a random email address.
// Returns: a random email address.
func (c *TRandomText) Email() string {
	return c.Words(2, 6) + "@" + c.Words(1, 3) + ".com"
}

// Generates a random text, consisting of first names, last names, colors, stuffs, adjectives, verbs, and punctuation marks.
// - minLength   minimum amount of words to generate. Text will contain "minSize" words if "maxSize" is omitted.
// - maxLength   (optional) maximum amount of words to generate.
// Returns:      a random text.
func (c *TRandomText) Text(minLength int, maxLength int) string {

	maxLength = ix.Max(minLength, maxLength)
	size := int(RandomInteger.NextInteger(int64(minLength), int64(maxLength)))

	result := ""
	result += RandomString.Pick(c.allWords)

	for len(result) < size {
		next := RandomString.Pick(c.allWords)
		if RandomBoolean.Chance(4, 6) {
			next = " " + strings.ToLower(next)
		} else if RandomBoolean.Chance(2, 5) {
			next = RandomString.PickChar(":,-") + strings.ToLower(next)
		} else if RandomBoolean.Chance(3, 5) {
			next = RandomString.PickChar(":,-") + " " + strings.ToLower(next)
		} else {
			next = RandomString.PickChar(".!?") + " " + next
		}

		result += next
	}

	return result
}
