package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		words := scanner.Text()
		if len(words) == 0 {
			continue
		}
		//commandName :=
		command, exists := getCommands()[words]
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

type commandCli struct {
	name        string
	description string
	callback    func() error
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
	}
}
