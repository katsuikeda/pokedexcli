package main

import (
	"fmt"
	"slices"
)

func commandHelp(cfg *config, args ...string) error {
	commands := getCommands()
	commandNames := make([]string, len(commands))

	i := 0
	for name := range commands {
		commandNames[i] = name
		i++
	}

	slices.Sort(commandNames)

	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	for _, name := range commandNames {
		fmt.Printf("%s: %s\n", commands[name].name, commands[name].description)
	}
	fmt.Println()
	return nil
}
