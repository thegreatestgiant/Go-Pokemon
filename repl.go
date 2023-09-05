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
	callback    func() error
}

func startRepl() {
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

		command.callback()
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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
