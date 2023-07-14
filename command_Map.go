package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error {

	//pokeApiClient := pokeapi.NewClient()

	resp, err := cfg.pokeApiClient.ListLocationAreas(cfg.nextLocationUrl)
	if err != nil {
		return err //log.Fatal(err)
	}

	fmt.Println("Location Areas :")

	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	cfg.nextLocationUrl = resp.Next
	cfg.previousLocationUrl = resp.Previous

	return nil
}
