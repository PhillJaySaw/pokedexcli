package main

import (
	"bufio"
	"fmt"
	"os"
)

type paginationConfig struct {
	nextLocationArea     string
	previousLocationArea *string
}

var Pagination paginationConfig

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()

		prompt := scanner.Text()
		command, exists := getCommands()[prompt]

		if !exists {
			fmt.Println("Command not found")
			continue
		}

		command.callback(&Pagination)
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*paginationConfig) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display 20 locations in the pokemon world. Each subsequent call will display the next 20 locations.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display 20 locations in the pokemon world. Each subsequent call will display the next 20 locations.",
			callback:    commandMapBack,
		},
	}
}
