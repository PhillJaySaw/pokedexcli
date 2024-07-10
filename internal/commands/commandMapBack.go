package commands

import (
	"fmt"
	pokeapi "github.com/philljaysaw/pokedexcli/internal/pokeApi"
)

func getPreviousLocationUrl(config *PaginationConfig) string {
	if config.previousLocationArea == nil {
		return "https://pokeapi.co/api/v2/location-area"
	}

	return *config.previousLocationArea
}
func MapBack(config *PaginationConfig) error {
	url := getPreviousLocationUrl(config)

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
