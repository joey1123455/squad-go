package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func Request(method, pUrl, key string, body io.Reader) (map[string]interface{}, error) {
	// Create a new URL
	u, err := url.Parse(pUrl)
	if err != nil {
		return nil, err
	}

	// Create a new request with the URL and method
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Parse the response body
	var resultMap map[string]interface{}
	err = json.Unmarshal(resBody, &resultMap)
	if err != nil {
		return nil, err
	}

	return resultMap, nil
}

func MakeRequest(body map[string]any, url, key, method string) (map[string]any, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return Request(method, url, key, bytes.NewBuffer(jsonBody))
}

// MakeGetRequest makes a GET request
func MakeGetRequest(queries map[string]string, pUrl, key string) (map[string]interface{}, error) {
	// Add queries to the URL
	if queries != nil {
		query := url.Values{}
		for key, value := range queries {
			query.Set(key, value)
		}
		pUrl += "?" + query.Encode()
	}

	return Request("GET", pUrl, key, nil)
}

// MakeDeleteRequest makes a DELETE request
func MakeDeleteRequest(pUrl, key string) (map[string]interface{}, error) {
	return Request("DELETE", pUrl, key, nil)
}
