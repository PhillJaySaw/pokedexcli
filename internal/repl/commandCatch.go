package repl

import (
	"errors"
	"fmt"
)

func Catch(config *Config, args []string) error {
	if len(args) == 0 {
		return errors.New("You need to provide a pokemon ID or name!")
	}

	pokemonID := args[0]

	pokemonDetails, err := config.Client.GetPokemon(pokemonID)

	if err != nil {
		return err
	}

	// TODO: using math/rand simulate catching the pokemon
	// chance should be based on BaseExperience
	// caught pokemons should be stored in the pokedex

	fmt.Println()
	fmt.Println("ID: ", pokemonDetails.ID)
	fmt.Println("NAME: ", pokemonDetails.Name)
	fmt.Println("BASE EXPERIENCE", pokemonDetails.BaseExperience)
	fmt.Println()

	return nil
}
