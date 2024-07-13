package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func (c *Client) GetPokemon(pokemonId string) (pokemonData PokemonResponse, error error) {
	url := fmt.Sprintf("%s/pokemon/%s", baseURL, pokemonId)

	res, err := c.httpClient.Get(url)

	if err != nil {
		error = err
		return
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	if res.StatusCode > 300 {
		fmt.Printf("Response failed with status code %d", res.StatusCode)
		error = errors.New("location request failed")
		return
	}

	json.Unmarshal(body, &pokemonData)

	return
}
