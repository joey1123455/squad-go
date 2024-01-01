package utils

import (
	"net/http"
	"net/http/httptest"
	"reflect"
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
	result, err := MakeRequest(body, url, apiKey, "post")

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
	result, err := MakeRequest(body, url, apiKey, "post")

	// Check assertions for error
	assert.Error(t, err)
	assert.Nil(t, result)
}

func Test_MakeGetRequest(t *testing.T) {
	// Mock the HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"key": "value"}`))
	}))
	defer server.Close()

	// Define the test cases
	tests := []struct {
		name          string
		queries       map[string]string
		url           string
		key           string
		wantResponses map[string]interface{}
		wantErr       bool
	}{
		{
			name:          "Valid Request",
			queries:       map[string]string{"param1": "value1"},
			url:           server.URL,
			key:           "key",
			wantResponses: map[string]interface{}{"key": "value"},
			wantErr:       false,
		},
		{
			name:          "No Queries",
			queries:       nil,
			url:           server.URL,
			key:           "key",
			wantResponses: map[string]interface{}{"key": "value"},
			wantErr:       false,
		},
		// More test cases...
	}

	// Run the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponses, err := MakeGetRequest(tt.queries, tt.url, tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeGetRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponses, tt.wantResponses) {
				t.Errorf("MakeGetRequest() = %v, want %v", gotResponses, tt.wantResponses)
			}
		})
	}
}
