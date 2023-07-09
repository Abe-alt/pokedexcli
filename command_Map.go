package main

import (
	"fmt"
	"github.com/Abe-alt/pokedexcli.git/internal/pokeapi"
	"log"
)

func commandMap() error {

	pokeApiClient := pokeapi.NewClient()

	resp, err := pokeApiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location Areas :")

	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	return nil
}
