package main

import (
	"fmt"
	"sort"
)

func commandHelp(cfg *config, args ...string) error {
	var commandSlice []cliCommand
	for _, ccom := range getCommands() {
		commandSlice = append(commandSlice, ccom)
	}
	sort.Slice(commandSlice, func(i int, j int) bool {
		return commandSlice[i].priotity < commandSlice[j].priotity
	})

	fmt.Println("Welcome To Pokemon")
	fmt.Println("Usage")
	for _, cmd := range commandSlice {
		fmt.Printf("  - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
