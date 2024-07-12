package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationArea(locationUrl *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if locationUrl != nil {
		url = *locationUrl
	}

	resp, err := c.httpClient.Get(url)
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
