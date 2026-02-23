package main

import "fmt"

func commandHelp(config ConfigCommand) error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, c := range commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}
