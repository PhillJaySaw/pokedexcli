package repl

import "os"

func Exit(config *Config, args []string) error {
	os.Exit(0)
	return nil
}
