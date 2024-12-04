package requests

import (
	"fmt"
	"io"
	"net/http"
)

func Delete(url string) (URLResponse, error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return URLResponse{}, fmt.Errorf("failed to create DELETE request to (%s): %v", url, err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return URLResponse{}, fmt.Errorf("failed to get a response from url (%s): %w", url, err)
	}
	defer resp.Body.Close()

	var bodyData []byte

	if resp.StatusCode != http.StatusNoContent {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return URLResponse{}, fmt.Errorf("failed to get read body from response (%s) : %w", url, err)
		}
		bodyData = body
	}

	return URLResponse{
		Body:       bodyData,
		StatusCode: resp.StatusCode,
	}, nil

}
