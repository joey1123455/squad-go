package utils

func ConvertToMap(anyValue interface{}) (map[string]interface{}, bool) {
	// Check if the value is a map with string keys and arbitrary values
	if value, ok := anyValue.(map[string]interface{}); ok {
		return value, true
	}

	// Return an empty map and false if the conversion is not successful
	return nil, false
}
