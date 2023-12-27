package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func MakeRequest(body map[string]any, url, key string) (map[string]any, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Create a new request with the JSON body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
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

	resBody, err := ioutil.ReadAll(response.Body)
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
