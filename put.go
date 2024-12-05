package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Put(url string, data interface{}) (URLResponse, error) {
	// Marshal the data to JSON
	jsonPayload, err := json.Marshal(data)
	if err != nil {
		return URLResponse{}, fmt.Errorf("error while marshalling data (%v)\n err: %w", data, err)
	}

	// Create a new HTTP request for PUT
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return URLResponse{}, fmt.Errorf("failed to create PUT request (%s): %w", url, err)
	}

	// Set the content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Execute the request using http.DefaultClient
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return URLResponse{}, fmt.Errorf("failed to get a response from url (%s): %w", url, err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return URLResponse{}, fmt.Errorf("failed to read body from response (%s): %w", url, err)
	}

	return URLResponse{
		Body:       body,
		StatusCode: resp.StatusCode,
	}, nil
}
