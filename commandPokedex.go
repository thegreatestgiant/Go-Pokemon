package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	cfg.theme.Info.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf("  - %s\n", cfg.themeFunc.Pokemon(pokemon.Name))
	}
	return nil
}
