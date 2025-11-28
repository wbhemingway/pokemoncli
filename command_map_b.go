package main

import (
	"fmt"
)

func commandMapB(cfg *config, args []string) error {
	if cfg.prev == nil {
		fmt.Println("You're on the first page")
		return nil
	}
	resp, err := cfg.pokeapiClient.ListLocations(cfg.prev)
	if err != nil {
		return err
	}
	cfg.next = resp.Next
	cfg.prev = resp.Prev
	for _, res := range resp.Results {
		fmt.Println(res.Name)
	}
	return nil
}
