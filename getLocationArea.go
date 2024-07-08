package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getLocationArea(url string) (LocationAreaResponse, error) {
	resp, err := http.Get(url)

	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode > 299 {
		fmt.Printf("Response failed with status code %d", resp.StatusCode)
	}

	var response LocationAreaResponse
	json.Unmarshal(body, &response)

	return response, nil
}
