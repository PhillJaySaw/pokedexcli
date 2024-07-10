package commands

import "os"

func Exit(config *PaginationConfig) error {
	os.Exit(0)
	return nil
}
