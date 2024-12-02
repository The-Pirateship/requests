package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Post(url string, data interface{}) ([]byte, int, error) {
	jsonPayload, err := json.Marshal(data)
	if err != nil {
		return nil, 0, fmt.Errorf("error while marshalling data (%v)\n err: %w", data, err)
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get a response from url (%s): %w", url, err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get read body from response (%s) : %w", url, err)
	}

	return body, response.StatusCode, nil
}
