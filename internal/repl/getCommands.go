package repl

type CliCommand struct {
	Callback    func(c *Config, args []string) error
	name        string
	description string
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Callback:    Help,
			name:        "help",
			description: "Display a help message",
		},
		"exit": {
			Callback:    Exit,
			name:        "exit",
			description: "Exit the Pokedex",
		},
		"map": {
			Callback:    Mapf,
			name:        "map",
			description: "Display 20 locations in the pokemon world. Each subsequent call will display the next 20 locations.",
		},
		"mapb": {
			Callback:    MapB,
			name:        "mapb",
			description: "Display 20 locations in the pokemon world. Each subsequent call will display the next 20 locations.",
		},
		"explore": {
			Callback:    Explore,
			name:        "explore",
			description: "Show list o Pokemon in a given area. Usage: explore {location name/id}",
		},
	}
}
