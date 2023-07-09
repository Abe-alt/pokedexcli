package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	for key, _ := range getCommands() {
		fmt.Printf("%s:  %s\n", key, getCommands()[key].description)
	}
	return nil
}
