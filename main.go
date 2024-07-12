package main

import (
	"time"

	pokeapi "github.com/philljaysaw/pokedexcli/internal/pokeApi"
	"github.com/philljaysaw/pokedexcli/internal/pokecache"
	"github.com/philljaysaw/pokedexcli/internal/repl"
)

func main() {
	cache := pokecache.NewCache(5 * time.Second)
	client := pokeapi.NewClient(10*time.Second, cache)
	config := &repl.Config{
		Client: client,
	}

	repl.StartRepl(config)
}
