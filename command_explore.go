package main

import (
	"fmt"
)

func commandExplore(c *config, param string) error {
	detailedArea, err := c.retr.ExploreLoaction(&param)
	if err != nil {
		return fmt.Errorf("Something went wrong")
	}
	for _, encounter := range detailedArea.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
