package repl

import (
	"fmt"
)

func Pokedex(config *Config, args []string) error {
	fmt.Println("Your Pokedex:")

	for id, pokemon := range config.PokeStore {
		normId := id

		for len(normId) < 3 {
			normId = "0" + normId
		}

		fmt.Printf("%s %s", normId, pokemon.Name)
		fmt.Println()
	}

	fmt.Println("---")

	return nil
}
