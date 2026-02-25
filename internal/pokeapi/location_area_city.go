package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) LocationAreaPerCity(cityName *string) (LocationAreaCityResp, error) {
	var locationArea LocationAreaCityResp

	if cityName == nil {
		return locationArea, fmt.Errorf("Does need to send an city name for the search")
	}

	url := baseURL + "/location-area/" + *cityName

	if cache, exists := c.cache.Get(url); exists {
		err := json.Unmarshal(cache, &locationArea)
		if err != nil {
			fmt.Println(err)
			return locationArea, err
		}
		return locationArea, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return locationArea, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return locationArea, err
	}
	if err != nil {
		return locationArea, err
	}

	c.cache.Add(url, body)

	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		fmt.Println(err)
		return locationArea, err
	}
	return locationArea, nil
}
