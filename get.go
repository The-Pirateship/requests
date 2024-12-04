package requests

import (
	"fmt"
	"io"
	"net/http"
)

func Get(url string) (URLResponse, error) {
	response, err := http.Get(url)
	if err != nil {
		return URLResponse{}, fmt.Errorf("failed to get a response from url (%s): %w", url, err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return URLResponse{}, fmt.Errorf("failed to get read body from response (%s) : %w", url, err)
	}

	return URLResponse{
		Body:       body,
		StatusCode: response.StatusCode,
	}, nil
}
