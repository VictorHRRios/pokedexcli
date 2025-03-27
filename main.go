package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		textBytes := scanner.Text()
		textString := strings.ToLower(string(textBytes))
		firstWord := strings.Fields(textString)[0]
		fmt.Println("Your command was:", firstWord)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
