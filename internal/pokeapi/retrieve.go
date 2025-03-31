package pokeapi

import (
	"github.com/VictorHRRios/pokedexcli/internal/pokecache"
	"time"
)

const (
	baseUrl = "https://pokeapi.co/api/v2"
)

type Retrieve struct {
	cache *pokecache.Cache
}

func GetRetrieve(timer int) *Retrieve {
	c := pokecache.NewCache(time.Duration(timer * int(time.Minute)))
	return &Retrieve{
		cache: &c,
	}
}
