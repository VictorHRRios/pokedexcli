package main

import "fmt"

func commandPokedex(c *config, s string) error {
	println("Your pokedex: ")
	for _, val := range c.pokedex {
		fmt.Printf("    - %v\n", val.Name)
	}

	return nil
}
