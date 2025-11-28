package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("no location to explore was given")
	}
	exploredArea, err := cfg.pokeapiClient.ListAreaPokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", exploredArea.Name)
	for _, pokemon := range exploredArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}