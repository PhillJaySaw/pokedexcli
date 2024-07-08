package main

import (
	"fmt"
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

func getNextLocationUrl(config *paginationConfig) string {
	if config.nextLocationArea == "" {
		return "https://pokeapi.co/api/v2/location-area"
	}

	return config.nextLocationArea
}

func commandMap(config *paginationConfig) error {
	url := getNextLocationUrl(config)

	locations, err := getLocationArea(url)

	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	config.nextLocationArea = locations.Next
	config.previousLocationArea = locations.Previous

	return nil
}
