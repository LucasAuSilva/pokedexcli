package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	id	 			 int
	name 			 string
	game_index int
} 

func commandMap(config ConfigCommand) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		fmt.Println("Some error has occur in the API, please try again later!")
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Println("Response failed with status code: %d", res.StatusCode)
		return err
	}
	if err != nil {
		fmt.Println("Some error has occur, please try again later!")
		return err
	}
	var locationsAreas []LocationArea
	err = json.Unmarshal(body, &locationsAreas)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(locationsAreas)
	return nil
}

