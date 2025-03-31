package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/VictorHRRios/pokedexcli/internal/pokeapi"
)

type config struct {
	retr *pokeapi.Retrieve
	next *string
	prev *string
}

func repl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	cfg.retr = pokeapi.GetRetrieve(5)
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
		command, ok := getCommands()[firstWord]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		command.callback(cfg)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
