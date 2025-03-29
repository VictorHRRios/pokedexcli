package main

import (
	"fmt"
	"github.com/VictorHRRios/pokedexcli/internal/pokeapi"
)

func commandMap(c *config) error {
	locations, err := pokeapi.ListLocations(c.next)
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

func commandMapB(c *config) error {
	locations, err := pokeapi.ListLocations(c.prev)
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
