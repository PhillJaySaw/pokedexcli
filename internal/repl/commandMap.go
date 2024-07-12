package repl

import (
	"fmt"
)

func getNextLocationUrl(config *Config) string {
	if config.nextLocationArea == nil {
		return "https://pokeapi.co/api/v2/location-area"
	}

	return *config.nextLocationArea
}

func getPreviousLocationUrl(config *Config) string {
	if config.previousLocationArea == nil {
		return "https://pokeapi.co/api/v2/location-area"
	}

	return *config.previousLocationArea
}

func Mapf(config *Config) error {
	url := getNextLocationUrl(config)

	locations, err := config.Client.GetLocationArea(&url)

	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	config.nextLocationArea = &locations.Next
	config.previousLocationArea = locations.Previous

	return nil
}

func MapB(config *Config) error {
	url := getPreviousLocationUrl(config)

	locations, err := config.Client.GetLocationArea(&url)

	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	config.nextLocationArea = &locations.Next
	config.previousLocationArea = locations.Previous

	return nil
}
