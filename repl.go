package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	c := &config{}
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
		command.callback(c)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
