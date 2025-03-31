package main

import "fmt"

func commandInspect(c *config, param string) error {
	pokemon, ok := c.pokedex[param]
	if !ok {
		return fmt.Errorf("Pokemon not caught!")
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, val := range pokemon.Stats {
		fmt.Printf("    - %v: %v\n", val.Stat.Name, val.BaseStat)
	}
	fmt.Println("Types:")
	for _, val := range pokemon.Types {
		fmt.Printf("    - %v\n", val.Type.Name)
	}

	return nil
}
