package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Thanks for using Pokedex!")
	os.Exit(0)
	return nil
}
