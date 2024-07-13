package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func (c *Client) GetLocationArea(locationName string) (locationAreaDetails LocationAreaResponse, error error) {
	url := fmt.Sprintf("%s/location-area/%s", baseURL, locationName)

	if cachedResponse, exists := c.cache.Get(url); exists {
		json.Unmarshal(cachedResponse, &locationAreaDetails)
		return
	}

	res, reqError := c.httpClient.Get(url)

	if reqError != nil {
		error = reqError
		return
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code %d", res.StatusCode)
		error = errors.New("location request failed")
		return
	}

	c.cache.Add(url, body)

	json.Unmarshal(body, &locationAreaDetails)

	return
}
