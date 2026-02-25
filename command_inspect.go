package main

import (
	"fmt"
)

func commandInspect(config *ConfigCommand) error {
	if len(config.parameters) < 2 {
			return fmt.Errorf("usage: inspect <pokemon>")
	}

	name := config.parameters[1]
	pokemon, exists := config.userPokemons[name]; if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	} 

	fmt.Println("Types:")
	for _, pType := range pokemon.Types {
		fmt.Printf("  - %s\n", pType.Type.Name)
	}

	return nil
}
