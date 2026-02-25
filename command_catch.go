package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *ConfigCommand) error {
	if len(config.parameters) < 2 {
			return fmt.Errorf("usage: catch <pokemon>")
	}

	pokeClient := config.pokeapiClient
	name := config.parameters[1]
	res, err := pokeClient.PokemonByName(&name)
	if err != nil {
		return fmt.Errorf("Error has occur in the API: %s", err.Error())
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	chance := min(max(70-res.BaseExperience/5, 5), 70)
	roll := rand.Intn(100)
	if roll < chance {
		fmt.Printf("%s was caught!\n", name)
		config.userPokemons[name] = res
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}
