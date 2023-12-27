package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConvertToMap_Success(t *testing.T) {
	// Test a successful conversion
	anyValue := map[string]interface{}{
		"key1": "value1",
		"key2": 42,
		"key3": true,
	}

	convertedMap, success := ConvertToMap(anyValue)

	// Assertions
	assert.True(t, success)
	assert.Equal(t, anyValue, convertedMap)
}

func Test_ConvertToMap_Failure(t *testing.T) {
	// Test a failure conversion
	anyValue := "not_a_map"

	convertedMap, success := ConvertToMap(anyValue)

	// Assertions
	assert.False(t, success)
	assert.Nil(t, convertedMap)
}
