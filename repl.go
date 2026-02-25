package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/LucasAuSilva/pokedexcli/internal/pokeapi"
)

func startRepl() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	config := ConfigCommand{
		pokeapiClient: *pokeapi.NewClient(),
	}
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
				config.parameters = cleanInputs
				err := command.callback(&config)
				if err != nil {
					fmt.Println(err.Error())
				}
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
	pokeapiClient pokeapi.Client
	parameters   	[]string
	NextURL 		 	*string
	PreviousURL  	*string
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
		"explore": {
			name: 			 "map",
			description: "Use to see all the pokemon on an given area",
			callback: 	 commandExplore,
		},
	}
}

