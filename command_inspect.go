package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("no pokemon to inspect was given")
	}
	pokemon, ok := cfg.pokedex[args[0]]
	if !ok {
		return errors.New("pokemon to be inspected has not been caught")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats{
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, types := range pokemon.Types {
		fmt.Printf(" - %s\n", types.Type.Name)
	}
	return nil
}