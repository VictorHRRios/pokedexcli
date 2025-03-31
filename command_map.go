package main

import (
	"fmt"
)

func commandMap(c *config, s string) error {
	locations, err := c.retr.ListLocations(c.next)
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

func commandMapB(c *config, s string) error {
	if c.prev == nil {
		fmt.Println("You are on the first page :)")
		return nil
	}

	locations, err := c.retr.ListLocations(c.prev)
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
