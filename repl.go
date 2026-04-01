package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/thegreatestgiant/Go-Pokemon/internal/commands"
	"github.com/thegreatestgiant/Go-Pokemon/internal/state"
)

func startRepl(cfg *commands.Config) {
	scanner := bufio.NewScanner(os.Stdin)
	cfg.Scanner = scanner
	for {
		fmt.Printf("%s", cfg.ThemeFunc.Prompt("Pokemon > "))

		scanner.Scan()
		text := parseInput(scanner.Text())
		if len(text) == 0 {
			continue
		}

		commandName := text[0]
		avaliableCommands := commands.GetCommands()

		command, ok := avaliableCommands[commandName]
		if !ok {
			cfg.Theme.Error.Printf("Enter a valid command\n")
			continue
		}

		args := []string{}
		if len(text) > 1 {
			args = text[1:]
		}

		err := command.Callback(cfg, args...)
		if err != nil {
			cfg.Theme.Error.Printf("error: %v\n", err)
		}

		saved := state.SavePokedex(cfg)
		if !saved && cfg.Debug {
			cfg.Theme.Error.Println("Didn't save pokedex")
		}

		saved = state.SaveSettings(cfg)
		if !saved && cfg.Debug {
			cfg.Theme.Error.Println("Didn't save settings")
		}
	}
}

func parseInput(str string) []string {
	lower := strings.ToLower(str)
	words := strings.Fields(lower)
	return words
}
