package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {

	pokemonName := args[0]
	pokemonData, err := cfg.pokeApiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	//Pokemon := pokemonData.Name
	//pokemonMap := make(map[string]string)

	r := rand.Intn(pokemonData.BaseExperience)
	fmt.Printf("Throwing the ball at `%s` , PokeBase: %v, RandNum: %v \n", pokemonData.Name, pokemonData.BaseExperience, r)

	if r > 20 {
		fmt.Printf("You caught `%s` \n", pokemonData.Name)
		cfg.caughtPoke[pokemonData.Name] = pokemonData

	} else {
		return fmt.Errorf("You missed catching %s ... \n Try again ! \n ", pokemonData.Name)
	}
	return nil
}
