package commands

import (
	"fmt"
)

func commandPokedex(cfg *Config, args ...string) error {
	cfg.Theme.Info.Println("Your Pokedex:")
	for _, pokemon := range cfg.Pokedex {
		fmt.Printf("  - %s\n", cfg.ThemeFunc.Pokemon(pokemon.Name))
	}
	return nil
}
