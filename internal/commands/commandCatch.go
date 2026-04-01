package commands

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *Config, args ...string) error {
	if len(args) < 1 {
		return errors.New("enter a {number/location} to explore")
	}

	pokemon := args[0]

	// TODO: make sure it's in the explore list

	pokemonUrl := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	pokemonStruct, err := cfg.PokeapiClient.GetPokemon(pokemonUrl)
	if err != nil {
		return err
	}

	baseXP := pokemonStruct.BaseExperience
	threshhold := 50
	randNum := rand.Intn(pokemonStruct.BaseExperience)
	if cfg.Debug {
		fmt.Println(baseXP, threshhold, randNum)
	}

	cfg.Theme.Info.Printf("Throwing a Pokeball at %s...", cfg.ThemeFunc.Pokemon(pokemon))
	time.Sleep(time.Second)
	if randNum > threshhold {
		return fmt.Errorf("failed to catch %s", cfg.ThemeFunc.Pokemon(pokemon))
	}

	cfg.Pokedex = append(cfg.Pokedex, pokemonStruct)
	cfg.Theme.Success.Printf("%s %s\n", cfg.ThemeFunc.Pokemon(pokemon), cfg.ThemeFunc.Success("was caught!!"))

	return nil
}
