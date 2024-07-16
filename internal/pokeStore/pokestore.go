package pokestore

import (
	"errors"

	pokeapi "github.com/philljaysaw/pokedexcli/internal/pokeApi"
)

type PokeStore struct {
	Store map[int]pokeapi.PokemonResponse
}

func (s *PokeStore) AddPokemon(data pokeapi.PokemonResponse) {
	s.Store[data.ID] = data
}

func (s *PokeStore) GetPokemon(ID int) (pokeapi.PokemonResponse, error) {
	pokemonData, ok := s.Store[ID]

	if !ok {
		return pokeapi.PokemonResponse{}, errors.New("Pokemon data not found")
	}

	return pokemonData, nil
}
