package main

import (
	"fmt"
)

func commandExplore(config *ConfigCommand) error {
	pokeClient := config.pokeapiClient
	res, err := pokeClient.LocationAreaPerCity(&config.parameters[1]); if err != nil {
		return fmt.Errorf("Error has occur in the API: %s", err.Error())
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", *encounter.Pokemon.Name)
	}
	return nil
}

