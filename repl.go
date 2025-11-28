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

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Gets Pokedex command information",
			callback:    commandHelp,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		val, ok := commands[words[0]]
		if !ok {
			fmt.Println("Unkown command")
		} else {
			val.callback()
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(text)))
}
