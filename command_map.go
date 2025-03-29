package main

import (
	"fmt"
	"github.com/VictorHRRios/pokedexcli/internal/pokeapi"
	"github.com/VictorHRRios/pokedexcli/internal/pokecache"
)

func commandMap(c *config, cache *pokecache.Cache) error {
	locations, err := pokeapi.ListLocations(c.next, cache)
	if err != nil {
		return fmt.Errorf("Something went wrong")
	}

	c.next = locations.Next
	c.prev = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapB(c *config, cache *pokecache.Cache) error {
	locations, err := pokeapi.ListLocations(c.prev, cache)
	if err != nil {
		return fmt.Errorf("Something went wrong")
	}
	if locations.Previous == nil {
		fmt.Println("You are on the first page :)")
		return nil
	}

	c.next = locations.Next
	c.prev = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}
