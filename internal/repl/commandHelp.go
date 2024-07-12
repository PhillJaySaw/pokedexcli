package repl

import "fmt"

func Help(config *Config) error {
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println()

	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}
