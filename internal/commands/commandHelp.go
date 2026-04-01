package commands

import (
	"fmt"
	"sort"
)

func commandHelp(cfg *Config, args ...string) error {
	var commandSlice []cliCommand
	for _, ccom := range GetCommands() {
		commandSlice = append(commandSlice, ccom)
	}
	sort.Slice(commandSlice, func(i int, j int) bool {
		return commandSlice[i].priotity < commandSlice[j].priotity
	})

	cfg.Theme.Header.Println("Welcome To Pokemon")
	cfg.Theme.Header.Println("Usage")
	for _, cmd := range commandSlice {
		fmt.Printf("  - %s%s: %s\n", cfg.ThemeFunc.Info(cmd.name), cfg.ThemeFunc.Highlight(cmd.argument), cmd.description)
	}
	fmt.Println()
	return nil
}
