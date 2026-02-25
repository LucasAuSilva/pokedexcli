package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) PokemonByName(pokemonName *string) (PokemonResp, error) {
	var pokemon PokemonResp

	if pokemonName == nil {
		return pokemon, fmt.Errorf("Does need to send an city name for the search")
	}

	url := baseURL + "/pokemon/" + *pokemonName

	if cache, exists := c.cache.Get(url); exists {
		err := json.Unmarshal(cache, &pokemon)
		if err != nil {
			return pokemon, err
		}
		return pokemon, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return pokemon, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return pokemon, err
	}
	if err != nil {
		return pokemon, err
	}

	c.cache.Add(url, body)

	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		fmt.Println(err)
		return pokemon, err
	}
	return pokemon, nil
}
