package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type LocationAreaListResp struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		ULR string `json:"url"`
	} `json:"results"`
} 

func (c *Client) GetLocationAreas(url *string) (LocationAreaListResp, error) {
	var res *http.Response
	var err error
	var locationsArea LocationAreaListResp
	if url == nil {
		res, err = c.httpClient.Get("https://pokeapi.co/api/v2/location-area")
	} else if !strings.Contains(*url, "location-area") {
		return locationsArea, fmt.Errorf("URL is not for location area")
	} else {
		res, err = c.httpClient.Get(*url)
	}
	if err != nil {
		fmt.Println("Some error has occur in the API, please try again later!")
		return locationsArea, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d", res.StatusCode)
		return locationsArea, err
	}
	if err != nil {
		fmt.Println("Some error has occur, please try again later!")
		return locationsArea, err
	}
	err = json.Unmarshal(body, &locationsArea)
	if err != nil {
		fmt.Println(err)
		return locationsArea, err
	}
	return locationsArea, nil
}
