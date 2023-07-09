package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		//words := strings.ToLower(scanner.Text())
		words := toLower(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue

		} else {
			fmt.Println("undefined command! Try Again ..")
			continue
		}

	}

}

func toLower(s string) []string {
	word := strings.ToLower(s)
	words := strings.Fields(word)
	return words
}

type commandCli struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	nextUrl     *string
	previousUrl *string
}

func getCommands() map[string]commandCli {
	return map[string]commandCli{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit the program",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "display maps",
			callback:    nil,
		},
	}
}
