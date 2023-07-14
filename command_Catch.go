package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {

	pokemonName := args[0]
	pokemon, err := cfg.pokeApiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	//Pokemon := pokemonData.Name
	//pokemonMap := make(map[string]string)

	r := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing the ball at `%s` , PokeBase: %v, RandNum: %v \n", pokemon.Name, pokemon.BaseExperience, r)

	if r > 20 {
		fmt.Printf("You caught `%s` \n", pokemon.Name)
		cfg.caughtPoke[pokemon.Name] = pokemon

	} else {
		return fmt.Errorf("You missed catching %s ... \n Try again ! \n ", pokemon.Name)
	}

	fmt.Println(" _______________________________________")
	fmt.Println("|         Pokemons you caught           |")
	fmt.Println("|_______________________________________|")

	for k, _ := range cfg.caughtPoke {
		fmt.Printf("- %s\n", k)
	}

	return nil
}
