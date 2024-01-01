package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func MakeRequest(body map[string]any, url, key, method string) (map[string]any, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Create a new request with the JSON body
	req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	// Set the Authorization header with the Bearer token
	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resultMap map[string]any
	err = json.Unmarshal(resBody, &resultMap)
	if err != nil {
		return nil, err
	}

	return resultMap, nil
}

func MakeGetRequest(queries map[string]string, pUrl, key string) (map[string]interface{}, error) {
	// Create a new URL with queries
	u, err := url.Parse(pUrl)
	if err != nil {
		return nil, err
	}

	if queries != nil {
		query := u.Query()
		for key, value := range queries {
			query.Set(key, value)
		}

		u.RawQuery = query.Encode()
	}

	// Create a new request with the updated URL
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	// Set the Authorization header with the Bearer token
	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resultMap map[string]interface{}
	err = json.Unmarshal(resBody, &resultMap)
	if err != nil {
		return nil, err
	}

	return resultMap, nil
}
