package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type LocationArea struct {
	id	 			 int
	name 			 string
	game_index int
} 

func (c *Client) getLocationAreas(url string) ([]LocationArea, error) {
	client := newClient()
	res, err := client.httpClient.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		fmt.Println("Some error has occur in the API, please try again later!")
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Println("Response failed with status code: %d", res.StatusCode)
		return nil, err
	}
	if err != nil {
		fmt.Println("Some error has occur, please try again later!")
		return nil, err
	}
	var locationsAreas []LocationArea
	err = json.Unmarshal(body, &locationsAreas)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(locationsAreas)
	return locationsAreas, nil
}
