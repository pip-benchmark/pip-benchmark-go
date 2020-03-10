package test_utilities

import (
	"testing"

	benchutil "github.com/pip-benchmark/pip-benchmark-go/utilities"
	"github.com/stretchr/testify/assert"
)

func TestProperties(t *testing.T) {
	//test load
	props := benchutil.Properties{}
	props.LoadFromFile("../../data/test.properties")

	assert.Equal(t, 4, len(props.Lines))
	assert.Equal(t, "", props.GetAsString("Key1", ""))
	assert.Equal(t, "Value2", props.GetAsString("Key2", ""))
	assert.Equal(t, `"Value 3"`, props.GetAsString("Key3", ""))

	//test save
	props = benchutil.Properties{}
	props.LoadFromFile("../../data/test.properties")
	props.SaveToFile("../../data/test_write.properties")

}
