package commands

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
)

func commandInspect(cfg *Config, args ...string) error {
	if len(args) < 1 {
		return errors.New("enter a pokemon to inspect")
	}

	pokemonName, flag, flagArgs := parseFlags(args...)
	moveFlags := []string{"m", "moves"}
	maxFlags := 4
	if flagArgs != nil {
		n, err := strconv.Atoi(flagArgs[0])
		if err != nil {
			if cfg.Debug {
				cfg.Theme.Error.Printf("Invalid flag argument: %s. Error: %s", flagArgs[0], err)
			}
			return err
		}
		maxFlags = n
	}

	pokemon, _, ok := cfg.findPokemon(pokemonName)
	if !ok {
		return fmt.Errorf("you have not caught %s, catch %s to inspect", pokemonName, pokemonName)
	}

	if cfg.Art {
		pokemon.PrintPokemonSprite()
	}

	fmt.Printf("%s %s\n", cfg.ThemeFunc.Info("Nickname:"), cfg.ThemeFunc.Pokemon(pokemon.NickName))
	fmt.Printf("%s %s\n", cfg.ThemeFunc.Info("Name:"), cfg.ThemeFunc.Pokemon(pokemon.Name))
	fmt.Printf("%s %v\n", cfg.ThemeFunc.Info("Height:"), cfg.ThemeFunc.Warning(pokemon.Height))
	fmt.Printf("%s %v\n", cfg.ThemeFunc.Info("Weight:"), cfg.ThemeFunc.Warning(pokemon.Weight))
	cfg.Theme.Info.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %v\n", cfg.ThemeFunc.Info(stat.Stat.Name), cfg.ThemeFunc.Warning(stat.BaseStat))
	}
	cfg.Theme.Info.Println("Types:")
	for _, val := range pokemon.Types {
		fmt.Printf("  - %s\n", cfg.ThemeFunc.Location(val.Type.Name))
	}
	if flag != "" && slices.Contains(moveFlags, flag) {
		cfg.Theme.Info.Println("Moves:")
		for i, move := range pokemon.Moves {
			if i >= maxFlags {
				break
			}
			fmt.Printf("  - %s\n", cfg.ThemeFunc.Warning(move.Move.Name))
		}
	}

	return nil
}
