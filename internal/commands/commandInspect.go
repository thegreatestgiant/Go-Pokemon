package commands

import (
	"errors"
	"fmt"

	"github.com/thegreatestgiant/Go-Pokemon/internal/pokeapi"
)

func commandInspect(cfg *Config, args ...string) error {
	if len(args) < 1 {
		return errors.New("enter a pokemon to inspect")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.findPokemon(pokemonName)
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

	return nil
}

func (cfg *Config) findPokemon(name string) (pokeapi.Pokemon, bool) {
	p, f := pokeapi.Pokemon{}, false
	for _, pokemon := range cfg.Pokedex {
		if pokemon.NickName == name {
			return pokemon, true
		} else if pokemon.Name == name {
			p = pokemon
			f = true
		}
	}
	return p, f
}
