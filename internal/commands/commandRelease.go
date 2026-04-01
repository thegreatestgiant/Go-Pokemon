package commands

import (
	"fmt"
	"strings"
)

func commandRelease(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a pokemon name to release")
	}
	nickname := strings.Join(args, " ")

	if nickname == "*" {
		loop := cfg.Pokedex
		for _, p := range loop {
			cfg.Theme.Success.Printf("Released %s!\n", p.NickName)
		}
		cfg.Pokedex = nil
	} else {
		if p, i, exists := cfg.findPokemon(nickname); exists {
			cfg.Pokedex = append(cfg.Pokedex[:i], cfg.Pokedex[1+i:]...)
			cfg.Theme.Success.Printf("Released %s!\n", p.NickName)
		} else {
			cfg.Theme.Warning.Printf("Could not release %s :(\n", nickname)
		}
	}

	return nil
}
