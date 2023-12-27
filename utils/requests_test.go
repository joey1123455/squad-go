package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MakeRequest_Successful(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

		// Simulate a successful response
		responseBody := `{"status": "success", "data": {"result": "ok"}}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responseBody))
	}))
	defer server.Close()

	// Set up test data
	body := map[string]interface{}{"key1": "value1", "key2": "value2"}
	url := server.URL
	apiKey := "test-token"

	// Call the function
	result, err := MakeRequest(body, url, apiKey)

	// Check assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "success", result["status"])
	assert.NotNil(t, result["data"])
	assert.Equal(t, "ok", result["data"].(map[string]interface{})["result"])
}

func TestMakeRequest_ErrorHandling(t *testing.T) {
	// Set up test data with an invalid URL
	body := map[string]interface{}{"key1": "value1", "key2": "value2"}
	url := "invalid_url"
	apiKey := "your_api_key"

	// Call the function
	result, err := MakeRequest(body, url, apiKey)

	// Check assertions for error
	assert.Error(t, err)
	assert.Nil(t, result)
}
