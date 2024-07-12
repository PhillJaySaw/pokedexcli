package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/philljaysaw/pokedexcli/internal/pokeApi"
	"github.com/philljaysaw/pokedexcli/internal/pokecache"
)

type Config struct {
	Client               pokeapi.Client
	Cache                *pokecache.Cache
	nextLocationArea     *string
	previousLocationArea *string
}

func StartRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()

		prompt := scanner.Text()
		command, exists := GetCommands()[prompt]

		if !exists {
			fmt.Println("Command not found")
			continue
		}

		command.Callback(cfg)
	}
}
