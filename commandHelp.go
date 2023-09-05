package main

import (
	"fmt"
)

func commandHelp(c *config) error {
	fmt.Println("Welcome To Pokemon")
	fmt.Println("Usage")
	for _, command := range getCommands() {
		fmt.Printf("  - %s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
