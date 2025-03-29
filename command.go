package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next string
	prev string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the name of 20 locations and advances a page",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the name of previous 20 locations and goes back a page",
			callback:    commandMapB,
		},
	}
}

type Command struct {
	Func func(*config) error
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage: \n\n")
	for _, val := range getCommands() {
		fmt.Printf("%v: %v\n", val.name, val.description)
	}
	return nil
}

func commandMap(c *config) error {
	var fullUrl string
	if c.next != "" {
		fullUrl = c.next
	} else {
		fullUrl = "https://pokeapi.co/api/v2/location-area/"
	}
	res, err := http.Get(fullUrl)
	if err != nil {
		return fmt.Errorf("Cannot fetch api")
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code %v", res.Status)
	}
	locationArea := LocationArea{}

	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return fmt.Errorf("Cannot unmarshal")
	}

	c.next = locationArea.Next
	c.prev = locationArea.Previous

	for _, location := range locationArea.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapB(c *config) error {
	var fullUrl string
	if c.next != "" {
		fullUrl = c.prev
	} else {
		fullUrl = "https://pokeapi.co/api/v2/location-area/"
	}
	res, err := http.Get(fullUrl)
	if err != nil {
		return fmt.Errorf("Cannot fetch api")
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code %v", res.Status)
	}
	locationArea := LocationArea{}

	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return fmt.Errorf("Cannot unmarshal")
	}

	c.next = locationArea.Next
	c.prev = locationArea.Previous

	for _, location := range locationArea.Results {
		fmt.Println(location.Name)
	}
	return nil
}
