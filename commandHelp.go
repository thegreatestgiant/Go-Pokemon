package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome To Pokemon")
	fmt.Println("Usage")
	for _, cmd := range getCommands() {
		fmt.Printf("  - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
