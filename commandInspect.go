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

	fmt.Printf("Name: %s\n", pokedex.Name)
	fmt.Printf("Height: %v\n", pokedex.Height)
	fmt.Printf("Weight: %v\n", pokedex.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokedex.Stats {
		fmt.Printf("  - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, val := range pokedex.Types {
		fmt.Printf("  - %s\n", val.Type.Name)
	}

	return nil
}
