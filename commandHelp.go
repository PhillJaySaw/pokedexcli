package main

import "fmt"

func commandHelp() error {
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}
