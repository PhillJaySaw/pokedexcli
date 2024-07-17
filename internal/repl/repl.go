package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/philljaysaw/pokedexcli/internal/pokeApi"
)

type Config struct {
	Client               pokeapi.Client
	PokeStore            map[string]pokeapi.PokemonResponse
	nextLocationArea     *string
	previousLocationArea *string
}

func StartRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()

		prompt := scanner.Text()
		args := strings.Split(prompt, " ")

		command, exists := GetCommands()[args[0]]

		if !exists {
			fmt.Println("Command not found")
			continue
		}

		error := command.Callback(cfg, args[1:])

		if error != nil {
			fmt.Println(error)
		}
	}
}
