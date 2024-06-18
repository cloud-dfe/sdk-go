package sdk_cloud_dfe

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type service struct {
	Config configService
}

func SetService(config configService) service {

	return service{
		Config: config,
	}
}

func (s service) Request(method, route string, payload interface{}) (interface{}, error) {
	// Create the URL
	url := fmt.Sprintf("%s:%d%s", s.Config.BaseURI, s.Config.Port, route)

	// Marshal the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// Create a new request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Config.Token))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client with timeout
	client := &http.Client{
		Timeout: time.Duration(s.Config.Timeout),
	}

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for non-200 status code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("received non-200 response: %d - %s", resp.StatusCode, string(bodyBytes)))
	}

	// Read the response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response body into a map
	var response interface{}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
