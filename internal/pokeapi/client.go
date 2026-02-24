package pokeapi

import (
	"net/http"
	"time"

	"github.com/LucasAuSilva/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient *http.Client
	cache *pokecache.Cache
}

func NewClient() *Client {
	cache := pokecache.NewCache(5 * time.Minute)
	return &Client{
		httpClient: &http.Client{},
		cache: cache,
	}
}
