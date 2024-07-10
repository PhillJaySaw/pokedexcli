package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/philljaysaw/pokedexcli/internal/commands"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()

		prompt := scanner.Text()
		command, exists := commands.GetCommands()[prompt]

		if !exists {
			fmt.Println("Command not found")
			continue
		}

		command.Callback(&commands.Pagination)
	}
}
