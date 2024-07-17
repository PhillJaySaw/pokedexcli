package repl

import (
	"errors"
	"fmt"
)

func Inspect(config *Config, args []string) error {

	if len(args) == 0 {
		return errors.New("Command is missing pokemon name!")
	}

	pokemonId := args[0]

	pokemon, ok := config.PokeStore[pokemonId]

	if !ok {
		fmt.Println("Pokemon not registered in pokedex.")
		return nil
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Height:", pokemon.Height)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%s: %v", stat.Stat.Name, stat.BaseStat)
		fmt.Println()
	}

	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("-%s", pokemonType.Type.Name)
		fmt.Println()
	}

	return nil
}
