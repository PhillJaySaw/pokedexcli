package commands

import "fmt"

func Help(config *PaginationConfig) error {
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println()

	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}
