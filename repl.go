package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	argument    string
	description string
	callback    func(*config, ...string) error
	priotity    int
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		intro := "Pokemon > "
		fmt.Printf("%s", ThemeFunc.Prompt(intro))

		scanner.Scan()
		text := parseInput(scanner.Text())
		if len(text) == 0 {
			continue
		}

		commandName := text[0]
		avaliableCommands := getCommands()

		command, ok := avaliableCommands[commandName]
		if !ok {
			Theme.Error.Printf("Enter a valid command")
			continue
		}

		args := []string{}
		if len(text) > 1 {
			args = text[1:]
		}

		err := command.callback(cfg, args...)
		if err != nil {
			Theme.Error.Printf("error: %v\n", err)
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
			priotity:    math.MinInt,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world.",
			callback:    commandMap,
			priotity:    10,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 locations in the Pokemon world.",
			callback:    commandMapb,
			priotity:    20,
		},
		"explore": {
			name:        "explore",
			argument:    " {(number of)/location}",
			description: "List of all the Pokémon in a given area\n    you can either enter the name of the location or the number",
			callback:    commandExplore,
			priotity:    30,
		},
		"catch": {
			name:        "catch",
			argument:    " {Pokemon}",
			description: "Catches Pokemon adds them to the user's Pokedex.",
			callback:    commandCatch,
			priotity:    40,
		},
		"inspect": {
			name:        "inspect",
			argument:    " {Pokemon}",
			description: "If the pokemon is in your pokedex then it will print it's stats",
			callback:    commandInspect,
			priotity:    50,
		},
		"pokedex": {
			name:        "pokedex",
			description: "print a list of all the names of the Pokemon you have caught.",
			callback:    commandPokedex,
			priotity:    60,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
			priotity:    math.MaxInt,
		},
	}
}
