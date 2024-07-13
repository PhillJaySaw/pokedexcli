package repl

import (
	"fmt"
	"slices"
)

func Help(config *Config, args []string) error {
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println()

	commands := GetCommands()

	keys := make([]string, 0, len(commands))
	for k := range commands {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	for _, key := range keys {
		cmd := commands[key]
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}
