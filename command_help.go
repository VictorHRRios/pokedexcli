package main

import (
	"fmt"

	"github.com/VictorHRRios/pokedexcli/internal/pokecache"
)

func commandHelp(c *config, cache *pokecache.Cache) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage: \n\n")
	for _, val := range getCommands() {
		fmt.Printf("%v: %v\n", val.name, val.description)
	}
	return nil
}
