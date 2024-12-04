package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Post(url string, data interface{}) (URLResponse, error) {
	jsonPayload, err := json.Marshal(data)
	if err != nil {
		return URLResponse{}, fmt.Errorf("error while marshalling data (%v)\n err: %w", data, err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return URLResponse{}, fmt.Errorf("failed to get a response from url (%s): %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return URLResponse{}, fmt.Errorf("failed to get read body from response (%s) : %w", url, err)
	}

	return URLResponse{
		Body:       body,
		StatusCode: resp.StatusCode,
	}, nil
}
