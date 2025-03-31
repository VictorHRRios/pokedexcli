package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(c *config, s string) error {

	pokemon, err := c.retr.GetPokemon(&s)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a ball at %v...\n", pokemon.Name)

	cathed := rand.IntN(400)
	catchProbability := min(pokemon.BaseExperience, 350)

	if cathed > catchProbability {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		c.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}
	return nil
}
