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
	callback    func(*config, ...string) error
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

		args := []string{}
		if len(text) > 1 {
			args = text[1:]
		}

		err := command.callback(cfg, args...)
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
		"explore": {
			name:        "explore {(number of)/location}",
			description: "list of all the Pokémon in a given area\n    you can either enter the name of the location or the number",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch {Pokemon}",
			description: "Catches Pokemon adds them to the user's Pokedex.",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
