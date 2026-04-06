package commands

import "math"

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
			priotity:    math.MinInt,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world.",
			Callback:    commandMap,
			priotity:    10,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 locations in the Pokemon world.",
			Callback:    commandMapb,
			priotity:    20,
		},
		"explore": {
			name:        "explore",
			argument:    "{(number of)/location}",
			description: "List of all the Pokémon in a given area\n    you can either enter the name of the location or the number",
			Callback:    commandExplore,
			priotity:    30,
		},
		"catch": {
			name:        "catch",
			argument:    "{Pokemon}",
			description: "Catches Pokemon adds them to the user's Pokedex.",
			Callback:    commandCatch,
			priotity:    40,
		},
		"release": {
			name:        "release",
			argument:    "{Pokemon}",
			description: "Releases Pokemon from your Pokedex. Use * to release all",
			Callback:    commandRelease,
			priotity:    42,
		},
		"inspect": {
			name:            "inspect",
			argument:        "{Pokemon}",
			flags:           "-m {X}, --moves {X}",
			flagDescription: "Will display the first X moves",
			description:     "If the pokemon is in your pokedex then it will print it's stats",
			Callback:        commandInspect,
			priotity:        50,
		},
		"pokedex": {
			name:        "pokedex",
			description: "print a list of all the names of the Pokemon you have caught.",
			Callback:    commandPokedex,
			priotity:    60,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    commandExit,
			priotity:    math.MaxInt,
		},
	}
}
