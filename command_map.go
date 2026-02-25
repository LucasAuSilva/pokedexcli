package main

import (
	"fmt"
)

func commandMap(config *ConfigCommand) error {
	pokeClient := config.pokeapiClient
	res, err := pokeClient.LocationAreaList(config.NextURL); if err != nil {
		return fmt.Errorf("Error has occur in the API: %s", err.Error())
	}
	for _, location := range res.Results {
		fmt.Println(location.Name)
	}
	config.NextURL = res.Next
	config.PreviousURL = res.Previous
	return nil
}

