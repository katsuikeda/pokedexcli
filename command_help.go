package main

import (
	"fmt"
	"slices"
)

func commandHelp() error {
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
		fmt.Printf("%s: %s\n", name, commands[name].description)
	}
	fmt.Println()
	return nil
}
