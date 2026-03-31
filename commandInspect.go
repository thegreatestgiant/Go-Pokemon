package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("enter a pokemon to inspect")
	}

	pokemon := args[0]

	pokedex, ok := cfg.pokedex[pokemon]
	if !ok {
		return fmt.Errorf("you have not caught %s, catch %s to inspect", pokemon, pokemon)
	}

	if cfg.art {
		pokedex.PrintPokemonSprite()
	}

	fmt.Printf("%s %s\n", cfg.themeFunc.Info("Name:"), cfg.themeFunc.Pokemon(pokedex.Name))
	fmt.Printf("%s %v\n", cfg.themeFunc.Info("Height:"), cfg.themeFunc.Warning(pokedex.Height))
	fmt.Printf("%s %v\n", cfg.themeFunc.Info("Weight:"), cfg.themeFunc.Warning(pokedex.Weight))
	cfg.theme.Info.Println("Stats:")
	for _, stat := range pokedex.Stats {
		fmt.Printf("  - %s: %v\n", cfg.themeFunc.Info(stat.Stat.Name), cfg.themeFunc.Warning(stat.BaseStat))
	}
	cfg.theme.Info.Println("Types:")
	for _, val := range pokedex.Types {
		fmt.Printf("  - %s\n", cfg.themeFunc.Location(val.Type.Name))
	}

	return nil
}
