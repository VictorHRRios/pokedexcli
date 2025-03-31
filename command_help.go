package main

import (
	"fmt"
)

func commandHelp(c *config, s string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage: \n\n")
	for _, val := range getCommands() {
		fmt.Printf("%v: %v\n", val.name, val.description)
	}
	return nil
}
