package utilities

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

type Properties struct {
	Lines      []*PropertyFileLine
	properties map[string]string
}

func (c *Properties) LoadFromFile(file string) {

	content, rdErr := ioutil.ReadFile(file)

	if rdErr != nil {
		panic("Can't read config file:" + file)
	}

	exp := regexp.MustCompile("/[\r\n]+/g")
	lines := exp.Split(string(content), -1)

	if c.Lines == nil {
		c.Lines = make([]*PropertyFileLine, 0)
	}

	for index := 0; index < len(lines); index++ {
		line := NewPropertyFileLine(lines[index], "", "")
		c.Lines = append(c.Lines, line)
	}

	if c.properties == nil {
		c.properties = make(map[string]string, 0)
	}

	c.populateItems()
}

func (c *Properties) populateItems() {
	for prop, _ := range c.properties {

		if _, ok := c.properties[prop]; ok && prop != "lines" {
			delete(c.properties, prop)
		}
	}

	for _, line := range c.Lines {
		if line.key != "" && len(line.key) > 0 {
			c.properties[line.key] = line.value
		}
	}
}

func (c *Properties) SaveToFile(file string) {
	c.synchronizeItems()

	content := ""
	for _, line := range c.Lines {
		content += fmt.Sprintln(line)
	}
	ioutil.WriteFile(file, []byte(content), 0755)
}

func (c *Properties) findLine(key string) *PropertyFileLine {
	for _, line := range c.Lines {
		if key == line.key {
			return line
		}
	}
	return nil
}

func (c *Properties) synchronizeItems() {
	// Update existing values and create missing lines
	for prop := range c.properties {
		//if (!c.hasOwnProperty(prop)) continue;

		if prop == "lines" {
			continue
		}

		line := c.findLine(prop)
		if line != nil {
			line.value = "" + c.properties[prop]
		} else {
			line = NewPropertyFileLine(prop, ""+c.properties[prop], "")
			c.Lines = append(c.Lines, line)
		}
	}

	// Remove lines mismatched with listed keys
	for index := len(c.Lines) - 1; index >= 0; index-- {
		line := c.Lines[index]
		if _, ok := c.properties[line.key]; line.key != "" && !ok {

			if index == len(c.Lines) {
				c.Lines = c.Lines[:index-1]
			} else {
				c.Lines = append(c.Lines[:index], c.Lines[index+1:]...)
			}
		}
	}
}

func (c *Properties) GetAsString(key string, defaultValue string) string {
	value := c.properties[key]
	if value == "" {
		return defaultValue
	}
	return value
}

func (c *Properties) SetAsString(key string, value string) {
	c.properties[key] = value
}

func (c *Properties) GetAsInteger(key string, defaultValue int) int {
	value := c.properties[key]
	if value == "" {
		return defaultValue
	}
	return Converter.StringToInteger(value, defaultValue)
}

func (c *Properties) SetAsInteger(key string, value int) {
	c.properties[key] = Converter.IntegerToString(value)
}

func (c *Properties) GetAsLong(key string, defaultValue int32) int32 {
	value := c.properties[key]
	if value == "" {
		return defaultValue
	}
	return Converter.StringToLong(value, defaultValue)
}

func (c *Properties) SetAsLong(key string, value int32) {
	c.properties[key] = Converter.LongToString(value)
}

func (c *Properties) GetAsDouble(key string, defaultValue float64) float64 {
	value := c.properties[key]
	if value == "" {
		return defaultValue
	}
	return Converter.StringToDouble(value, defaultValue)
}

func (c *Properties) SetAsDouble(key string, value float64) {
	c.properties[key] = Converter.DoubleToString(value)
}

func (c *Properties) GetAsBoolean(key string, defaultValue bool) bool {
	value := c.properties[key]
	if value == "" {
		return defaultValue
	}
	return Converter.StringToBoolean(value, defaultValue)
}

func (c *Properties) SetAsBoolean(key string, value bool) {
	c.properties[key] = Converter.BooleanToString(value)
}
