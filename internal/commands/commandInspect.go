package commands

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *Config, args ...string) error {
	if len(args) < 1 {
		return errors.New("enter a pokemon to inspect")
	}

	pokemon := args[0]

	pokedex, ok := cfg.Pokedex[pokemon]
	if !ok {
		return fmt.Errorf("you have not caught %s, catch %s to inspect", pokemon, pokemon)
	}

	if cfg.Art {
		pokedex.PrintPokemonSprite()
	}

	fmt.Printf("%s %s\n", cfg.ThemeFunc.Info("Name:"), cfg.ThemeFunc.Pokemon(pokedex.Name))
	fmt.Printf("%s %v\n", cfg.ThemeFunc.Info("Height:"), cfg.ThemeFunc.Warning(pokedex.Height))
	fmt.Printf("%s %v\n", cfg.ThemeFunc.Info("Weight:"), cfg.ThemeFunc.Warning(pokedex.Weight))
	cfg.Theme.Info.Println("Stats:")
	for _, stat := range pokedex.Stats {
		fmt.Printf("  - %s: %v\n", cfg.ThemeFunc.Info(stat.Stat.Name), cfg.ThemeFunc.Warning(stat.BaseStat))
	}
	cfg.Theme.Info.Println("Types:")
	for _, val := range pokedex.Types {
		fmt.Printf("  - %s\n", cfg.ThemeFunc.Location(val.Type.Name))
	}

	return nil
}
