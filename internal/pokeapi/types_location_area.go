package pokeapi

type LocationAreaListResp struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name 							string `json:"name"`
		ULR 							string `json:"url"`
	} `json:"results"`
}

type LocationAreaCityResp struct {
	Id    						*int 		`json:"id"`
	Name  						*string `json:"name"`
	PokemonEncounters []PokemonEncountersResp `json:"pokemon_encounters"`
}

type PokemonEncountersResp struct {
	Pokemon PokemonResp `json:"pokemon"` 
}

type PokemonResp struct {
	Name *string `json:"name"`
	Url  *string `json:"url"`
}
