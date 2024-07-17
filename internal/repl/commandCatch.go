package repl

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

const (
	MinBaseExperience = 20.00
	MaxBaseExperience = 500.00
)

func calculateCatchRate(baseExperience float64) float64 {
	div := MaxBaseExperience - MinBaseExperience
	catchRate := 100 * (1 - ((baseExperience - MinBaseExperience) / div))
	return catchRate
}

func Catch(config *Config, args []string) error {
	if len(args) == 0 {
		return errors.New("You need to provide a pokemon ID or name!")
	}

	pokemonID := args[0]

	pokemonDetails, err := config.Client.GetPokemon(pokemonID)

	if err != nil {
		return err
	}

	baseExp := pokemonDetails.BaseExperience
	catchRate := calculateCatchRate(float64(baseExp))
	randFloat := float64(rand.IntN(100))

	name := pokemonDetails.Name

	fmt.Println("Pokemon:", name)
	fmt.Println("Base experience", baseExp)
	fmt.Println("Catch rate: ", catchRate)

	if randFloat < catchRate {
		fmt.Printf("You caught %s!!!", name)
		config.PokeStore[pokemonID] = pokemonDetails
		fmt.Println()
	} else {
		fmt.Printf("%s got away...", name)
		fmt.Println()
	}

	return nil
}
