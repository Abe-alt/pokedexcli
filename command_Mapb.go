package main

import (
	"fmt"
	"log"
)

func commandMapb(cfg *config) error {

	//pokeApiClient := pokeapi.NewClient()

	resp, err := cfg.pokeApiClient.ListLocationAreas(cfg.previousLocationUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location Areas :")

	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	cfg.nextLocationUrl = resp.Next
	cfg.previousLocationUrl = resp.Previous

	return nil
}
