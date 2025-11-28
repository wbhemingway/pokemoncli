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
	callback    func(*config) error
}

type config struct {
	next *string
	prev *string
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
		"map": {
			name:        "map",
			description: "Gets the next 20 location areas",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Gets the last 20 location areas",
			callback:    commandMapB,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := config{}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		val, ok := commands[words[0]]
		if !ok {
			fmt.Println("Unkown command")
		} else {
			err := val.callback(&cfg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(text)))
}
