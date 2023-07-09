package main

import (
	"fmt"
	"github.com/Abe-alt/pokedexcli.git/internal/pokeapi"
	"log"
)

func main() {

	pokeApiClient := pokeapi.NewClient()

	resp, err := pokeApiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

	//startRepl()
}
