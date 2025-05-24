package pokeapi

import (
	"net/http"
	"time"

	"github.com/bdbrwr/bootdev_pokedex/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
}

// NewClient -
func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
