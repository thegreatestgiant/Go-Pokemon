package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokemon >")

		scanner.Scan()
		text := parseInput(scanner.Text())
		if len(text) == 0 {
			continue
		}

		commandName := text[0]
		avaliableCommands := getCommands()

		command, ok := avaliableCommands[commandName]
		if !ok {
			fmt.Println("Enter a valid command")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}
}

func parseInput(str string) []string {
	lower := strings.ToLower(str)
	words := strings.Fields(lower)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the names of the previous 20 locations in the Pokemon world.",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
