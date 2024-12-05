package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("Thanks for using Pokedex!")
	os.Exit(0)
	return nil
}
