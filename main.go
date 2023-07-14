package main

import (
	"github.com/Abe-alt/pokedexcli.git/internal/pokeapi"
)

type config struct {
	pokeApiClient       pokeapi.Client // to keep using http client more efficiently
	nextLocationUrl     *string        // if its nil then we don't have a next or prev page
	previousLocationUrl *string
	caughtPoke          map[string]pokeapi.Pokemons
}

func main() {
	cfg := config{
		pokeApiClient: pokeapi.NewClient(),
		caughtPoke:    make(map[string]pokeapi.Pokemons),
	}

	startRepl(&cfg)
}
