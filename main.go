package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		textBytes := scanner.Text()
		textString := strings.ToLower(string(textBytes))
		firstWord := strings.Fields(textString)[0]
		command, ok := getCommands()[firstWord]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		command.callback()
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage: \n\n")
	for _, val := range getCommands() {
		fmt.Printf("%v: %v\n", val.name, val.description)
	}
	return nil
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
