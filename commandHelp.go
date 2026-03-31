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

	cfg.theme.Header.Println("Welcome To Pokemon")
	cfg.theme.Header.Println("Usage")
	for _, cmd := range commandSlice {
		fmt.Printf("  - %s%s: %s\n", cfg.themeFunc.Info(cmd.name), cfg.themeFunc.Highlight(cmd.argument), cmd.description)
	}
	fmt.Println()
	return nil
}
