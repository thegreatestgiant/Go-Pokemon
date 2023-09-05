package main

import (
	"errors"
	"fmt"
	"math/rand"
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

	fmt.Printf("Throwing a Pokeball at %s...", pokemon)
	if randNum > threshhold {
		return fmt.Errorf("failed to catch %s", pokemon)
	}

	cfg.pokedex[pokemon] = pokemonStruct
	fmt.Printf("%s was caught!!\n", pokemon)

	return nil
}
