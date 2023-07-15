package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {

	fmt.Println("Your Pokedex:")
	for pokemon, _ := range cfg.caughtPoke {
		fmt.Printf(" -%s \n", pokemon)
	}
	return nil

}
