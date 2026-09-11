package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	var commands map[string]cliCommand = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "displays the next 20 location areas in the Pokemon world ",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the previous 20 location areas in the Pokemon world ",
			callback:    commandMapb,
		},
	}
	return commands
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	commands := config.commands
	for _, value := range commands {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	return nil
}
