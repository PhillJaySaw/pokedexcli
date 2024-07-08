package main

import "fmt"

func getPreviousLocationUrl(config *paginationConfig) string {
	if config.previousLocationArea == nil {
		return "https://pokeapi.co/api/v2/location-area"
	}

	return *config.previousLocationArea
}
func commandMapBack(config *paginationConfig) error {
	url := getPreviousLocationUrl(config)

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
