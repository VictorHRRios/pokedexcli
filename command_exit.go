package main

import (
	"fmt"
	"os"

	"github.com/VictorHRRios/pokedexcli/internal/pokecache"
)

func commandExit(c *config, cache *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
