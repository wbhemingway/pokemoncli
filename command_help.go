package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for command := range commands {
		fmt.Printf("%v: %v\n", commands[command].name, commands[command].description)
	}

	return nil
}
