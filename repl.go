package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/VictorHRRios/pokedexcli/internal/pokeapi"
)

type config struct {
	pokedex map[string]pokeapi.PokemonDetail
	retr    *pokeapi.Retrieve
	next    *string
	prev    *string
}

func repl(cfg *config) {
	var param string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		textBytes := scanner.Text()
		cleanText := cleanInput(string(textBytes))
		if len(cleanText) == 0 {
			fmt.Println("No provided command")
			continue
		}
		firstWord := cleanText[0]

		if len(cleanText) > 1 {
			param = cleanText[1]
		}
		command, ok := getCommands()[firstWord]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg, param)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
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
		"explore": {
			name:        "explore",
			description: "Displays the names of the pokemon in a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Displays a toy catch command.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays the details of a caught pokemon.",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays the caught pokemon",
			callback:    commandPokedex,
		},
	}
}
