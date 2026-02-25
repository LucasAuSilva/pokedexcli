package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) LocationAreaList(pageURL *string) (LocationAreaListResp, error) {
	var locationsArea LocationAreaListResp
	url := baseURL + "/location-area"

	if pageURL != nil && *pageURL != "" {
		url = *pageURL
	}

	if cache, exists := c.cache.Get(url); exists {
		err := json.Unmarshal(cache, &locationsArea)
		if err != nil {
			return locationsArea, err
		}
		return locationsArea, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return locationsArea, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return locationsArea, err
	}
	if err != nil {
		return locationsArea, err
	}

	c.cache.Add(url, body)

	err = json.Unmarshal(body, &locationsArea)
	if err != nil {
		fmt.Println(err)
		return locationsArea, err
	}
	return locationsArea, nil
}
