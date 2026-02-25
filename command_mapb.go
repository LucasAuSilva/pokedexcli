package main

import (
	"fmt"

	"github.com/LucasAuSilva/pokedexcli/internal/pokeapi"
)

func commandMapb(config *ConfigCommand) error {
	pokeClient := pokeapi.NewClient()
	res, err := pokeClient.LocationAreaList(config.PreviousURL); if err != nil {
		return fmt.Errorf("Error has occur in the API: %s", err.Error())
	}
	for _, location := range res.Results {
		fmt.Println(location.Name)
	}
	config.NextURL = res.Next
	config.PreviousURL = res.Previous
	return nil
}
