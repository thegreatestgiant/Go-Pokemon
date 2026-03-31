package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("enter a {number/location} to explore")
	}

	pokemon := args[0]
	pokemonUrl := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	pokemonStruct, err := cfg.pokeapiClient.GetPokemon(pokemonUrl)
	if err != nil {
		return err
	}

	baseXP := pokemonStruct.BaseExperience
	threshhold := 50
	randNum := rand.Intn(pokemonStruct.BaseExperience)
	fmt.Println(baseXP, threshhold, randNum)

	cfg.theme.Info.Printf("Throwing a Pokeball at %s...", cfg.themeFunc.Pokemon(pokemon))
	time.Sleep(time.Second)
	if randNum > threshhold {
		return fmt.Errorf("failed to catch %s", cfg.themeFunc.Pokemon(pokemon))
	}

	cfg.pokedex[pokemon] = pokemonStruct
	cfg.theme.Success.Printf("%s %s\n", cfg.themeFunc.Pokemon(pokemon), cfg.themeFunc.Success("was caught!!"))

	return nil
}
