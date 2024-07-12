package main

import (
	"time"

	pokeapi "github.com/philljaysaw/pokedexcli/internal/pokeApi"
	"github.com/philljaysaw/pokedexcli/internal/repl"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	config := &repl.Config{
		Client: client,
	}

	repl.StartRepl(config)
}
