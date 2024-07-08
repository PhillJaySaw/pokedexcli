package main

import "os"

func commandExit(config *paginationConfig) error {
	os.Exit(0)
	return nil
}
