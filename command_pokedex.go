package main

import (
	"fmt"
)

func commandPokedex(config *ConfigCommand) error {
	if len(config.userPokemons) < 1 {
		fmt.Println("You didn't caught any pokemon yet!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.userPokemons {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
