package main

import (
	"github.com/wbhemingway/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	pokedex := make(map[string]pokeapi.RespPokemon)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex: pokedex,
	}
	startRepl(cfg)
}
