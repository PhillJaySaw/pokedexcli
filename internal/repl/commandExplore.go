package repl

import (
	"errors"
	"fmt"
)

func Explore(config *Config, args []string) error {
	if len(args) == 0 {
		return errors.New("No location name provided")
	}

	locationName := args[0]

	locationDetails, error := config.Client.GetLocationArea(locationName)

	if error != nil {
		return error
	}

	for _, pokemon := range locationDetails.PokemonEncounters {
		id, _ := getLocationId(pokemon.Pokemon.URL)
		fmt.Println(pokemon.Pokemon.Name, id)
	}

	return nil
}
