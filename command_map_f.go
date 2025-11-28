package main

import (
	"fmt"
)

func commandMapF(cfg *config, args []string) error {
	resp, err := cfg.pokeapiClient.ListLocations(cfg.next)
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
