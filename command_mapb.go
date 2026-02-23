package main

import (
	"fmt"

	"github.com/LucasAuSilva/pokedexcli/internal/pokeapi"
)

func commandMapb(config *ConfigCommand) error {
	pokeClient := pokeapi.NewClient()
	res, err := pokeClient.GetLocationAreas(config.Previous); if err != nil {
		return fmt.Errorf("Error has occur in the API: %s", err.Error())
	}
	for _, location := range res.Results {
		fmt.Println(location.Name)
	}
	config.Next = res.Next
	config.Previous = res.Previous
	return nil
}
