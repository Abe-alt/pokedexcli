package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("wrong args")
	}
	locationName := args[0]
	fmt.Println(locationName)
	resp, err := cfg.pokeApiClient.GetLocationArea(locationName)
	if err != nil {
		return err
	}
	fmt.Println("Pokemon IDs :")
	for _, id := range resp.PokemonEncounters {
		fmt.Printf("- %s\n ", id.Pokemon.Name)
	}
	return nil
}
