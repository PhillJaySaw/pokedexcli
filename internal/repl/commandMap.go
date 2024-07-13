package repl

import (
	"fmt"
	"net/url"
	"path"
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

func getLocationId(locationUrl string) (string, error) {
	parsedURL, error := url.Parse(locationUrl)

	if error != nil {
		return "", error
	}

	id := path.Base(parsedURL.Path)
	return id, nil
}

func Mapf(config *Config, args []string) error {
	url := getNextLocationUrl(config)

	locations, err := config.Client.GetLocationAreas(&url)

	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		locationId, _ := getLocationId(location.URL)
		fmt.Println(locationId, location.Name)
	}

	config.nextLocationArea = &locations.Next
	config.previousLocationArea = locations.Previous

	return nil
}

func MapB(config *Config, args []string) error {
	url := getPreviousLocationUrl(config)

	locations, err := config.Client.GetLocationAreas(&url)

	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		locationId, _ := getLocationId(location.URL)
		fmt.Println(locationId, location.Name)
	}

	config.nextLocationArea = &locations.Next
	config.previousLocationArea = locations.Previous

	return nil
}
