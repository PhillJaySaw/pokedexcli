package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func (c *Client) GetLocationAreas(locationUrl *string) (res LocationAreasResponse, err error) {
	url := baseURL + "/location-area"
	if locationUrl != nil {
		url = *locationUrl
	}

	if cachedResponse, ok := c.cache.Get(url); ok {
		json.Unmarshal(cachedResponse, &res)

		return
	}

	resp, reqError := c.httpClient.Get(url)
	if reqError != nil {
		err = reqError
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode > 299 {
		fmt.Printf("Response failed with status code %d", resp.StatusCode)
		err = errors.New("location request failed")
		return
	}

	c.cache.Add(url, body)

	json.Unmarshal(body, &res)

	return
}
