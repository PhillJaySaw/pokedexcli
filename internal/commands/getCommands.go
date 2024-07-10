package commands

type PaginationConfig struct {
	nextLocationArea     string
	previousLocationArea *string
}

type CliCommand struct {
	Callback    func(*PaginationConfig) error
	name        string
	description string
}

var Pagination PaginationConfig

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
			Callback:    Map,
			name:        "map",
			description: "Display 20 locations in the pokemon world. Each subsequent call will display the next 20 locations.",
		},
		"mapb": {
			Callback:    MapBack,
			name:        "mapb",
			description: "Display 20 locations in the pokemon world. Each subsequent call will display the next 20 locations.",
		},
	}
}
