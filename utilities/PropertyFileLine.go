package utilities

import (
	"regexp"
	"strings"
)

type PropertyFileLine struct {
	line    string
	key     string
	value   string
	comment string
}

func NewPropertyFileLine(key string, value string, comment string) *PropertyFileLine {
	c := PropertyFileLine{}
	if value == "" && comment == "" {
		c.parseLine(key)
	} else {
		c.key = key
		c.value = value
		c.comment = comment
		c.composeNewLine()
	}
	return &c
}

func (c *PropertyFileLine) Key() string {
	return c.key
}

func (c *PropertyFileLine) GetValue() string {
	return c.value
}

func (c *PropertyFileLine) SetValue(value string) {
	c.value = value
	c.composeNewLine()
}

func (c *PropertyFileLine) GetComment() string {
	return c.comment
}

func (c *PropertyFileLine) SetComment(value string) {
	c.comment = value
	c.composeNewLine()
}

func (c *PropertyFileLine) Line() string {
	return c.line
}

func (c *PropertyFileLine) composeNewLine() {
	c.line = ""
	if c.key != "" && len(c.key) > 0 {
		c.line += c.encodeValue(c.key)
		c.line += "="
		c.line += c.encodeValue(c.value)
	}
	if c.comment != "" && len(c.comment) > 0 {
		c.line += " ;"
		c.line += c.comment
	}
}

func (c *PropertyFileLine) parseLine(line string) {
	c.line = line

	// Parse comment
	commentIndex := c.indexOfComment(line)
	if commentIndex >= 0 {
		c.comment = line[(commentIndex + 1):]
		line = line[:commentIndex]
	}

	// Parse key and value
	assignmentIndex := strings.Index(line, "=")
	if assignmentIndex >= 0 {
		c.value = line[(assignmentIndex + 1):]
		c.value = c.decodeValue(c.value)
		c.key = line[:assignmentIndex]
		c.key = c.decodeValue(c.key)
	} else {
		c.key = c.decodeValue(line)
		c.value = ""
	}
}

func (c *PropertyFileLine) indexOfComment(value string) int {
	partOfString := false
	var stringDelimiter byte = ' '
	for index := 0; index < len(value); index++ {
		chr := value[index]
		if partOfString == false && chr == ';' {
			return index
		} else if partOfString == true && chr == stringDelimiter {
			partOfString = false
		} else if partOfString == false && chr == '"' || chr == '\'' {
			partOfString = true
			stringDelimiter = chr
		}
	}
	return -1
}

func (c *PropertyFileLine) decodeValue(value string) string {

	value = strings.TrimSpace(value)
	if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
		value = value[1 : len(value)-2]

		exp := regexp.MustCompile("/\"\"/g")
		value = exp.ReplaceAllString(value, "\"")
	}
	if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") {
		value = value[1 : len(value)-1]

		exp := regexp.MustCompile("/''/g")
		value = exp.ReplaceAllString(value, "'")
	}
	return value
}

func (c *PropertyFileLine) encodeValue(value string) string {
	if value == "" {
		return value
	}

	if strings.HasPrefix(value, " ") || strings.HasSuffix(value, " ") || strings.Index(value, ";") >= 0 {
		exp := regexp.MustCompile("/\"/g")
		value = exp.ReplaceAllString(value, "\"\"")
		value = "\"" + value + "\""
	}
	return value
}

func (c *PropertyFileLine) ToString() string {
	return c.line
}
