package main

import (
	"fmt"
)

func commandHelp(c *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for command := range commands {
		fmt.Printf("%v: %v\n", commands[command].name, commands[command].description)
	}

	return nil
}
