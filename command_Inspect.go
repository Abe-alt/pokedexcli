package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("wrong args")
	}
	pokemonName := args[0]

	for k, v := range cfg.caughtPoke {
		if k != pokemonName {
			fmt.Println("you have not caught that pokemon !")
		}
		if k == pokemonName {
			fmt.Printf(" Name: %s\n Height: %v\n Weight: %v\n", v.Name, v.Height, v.Weight)
			fmt.Println("Types: ")
			for _, types := range v.Types {
				fmt.Printf(" -%s\n", types.Type.Name)
			}
			fmt.Println("Stats: ")
			for _, stats := range v.Stats {
				fmt.Printf(" -%s\n", stats.Stat.Name)
			}
		}
	}
	return nil
}
