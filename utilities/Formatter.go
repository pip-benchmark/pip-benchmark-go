package utilities

import (
	"fmt"
	"strconv"
	"time"
)

var Formatter TFormatter = TFormatter{}

type TFormatter struct {
}

func (c *TFormatter) PadLeft(value string, length int, padSymbol string) string {
	output := ""
	output += padSymbol
	output += value
	output += padSymbol

	for len(output) < length+2 {
		output = padSymbol + output
	}
	return output
}

func (c *TFormatter) PadRight(value string, length int, padSymbol string) string {
	output := ""
	output += padSymbol
	output += value
	output += padSymbol

	for len(output) < length+2 {
		output += padSymbol
	}
	return output
}

func (c *TFormatter) FormatNumber(value float64, decimals int) string {
	return strconv.FormatFloat(value, 'f', decimals, 64)
}

func (c *TFormatter) FormatDate(date time.Time) string {
	return date.UTC().Format("2006-01-02")
}

func (c *TFormatter) FormatTime(date time.Time) string {
	return date.UTC().Format("15:04:05")
}

func (c *TFormatter) FormatTimeSpan(ticks int64) string {
	millis := (ticks % 1000)
	seconds := ((ticks / 1000) % 60)
	minutes := ((ticks / 1000 / 60) % 60)
	hours := (ticks / 1000 / 60 / 60)
	return fmt.Sprintf("%d:%d:%d.%d", hours, minutes, seconds, millis)
}
