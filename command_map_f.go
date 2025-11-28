package main

import (
	"fmt"
	"github.com/wbhemingway/pokedexcli/internal/pokeapi"
)

func commandMapF(cfg *config) error {
	names, newPrev, newNext, err := pokeapi.GetLocationAreas(cfg.next)
	if err != nil {
		return err
	}
	cfg.next = newNext
	cfg.prev = newPrev
	for _, name := range names {
		fmt.Println(name)
	}
	return nil
}
