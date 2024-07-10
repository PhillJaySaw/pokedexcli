package commands

import (
	"fmt"
	pokeapi "github.com/philljaysaw/pokedexcli/internal/pokeApi"
)

func getNextLocationUrl(config *PaginationConfig) string {
	if config.nextLocationArea == "" {
		return "https://pokeapi.co/api/v2/location-area"
	}

	return config.nextLocationArea
}

func Map(config *PaginationConfig) error {
	url := getNextLocationUrl(config)

	locations, err := pokeapi.GetLocationArea(url)

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
