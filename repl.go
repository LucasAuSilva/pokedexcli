package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	config := ConfigCommand{}
	for {
		fmt.Print("Pokedex > ")
		res := scanner.Scan()
		if res {
			input := scanner.Text()
			if len(input) == 0 {
				fmt.Printf("Your command was: %s\n", input)
				continue
			}
			cleanInputs := cleanInput(input)
			if command, exists := commands[cleanInputs[0]]; exists {
				command.callback(&config)
			} else {
				fmt.Println("Unknown command.")
			}
		}
	}
}

func cleanInput(text string) []string {
	cleanInputs := strings.Fields(
			strings.ToLower(text))
	return cleanInputs
}

type ConfigCommand struct {
	Next 		 *string
	Previous *string
}

type CliCommand struct {
	name 				string
	description string
	callback 		func(config *ConfigCommand) error
}

func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			name: 			 "exit",
			description: "Exit the Pokedex",
			callback: 	 commandExit,
		},
		"help": {
			name: 			 "help",
			description: "Displays a help message",
			callback: 	 commandHelp,
		},
		"map": {
			name: 			 "map",
			description: "Display the next 20 locations areas of Pokemon World",
			callback: 	 commandMap,
		},
		"mapb": {
			name: 			 "map",
			description: "Display the previous 20 locations areas of Pokemon World",
			callback: 	 commandMapb,
		},
	}
}

