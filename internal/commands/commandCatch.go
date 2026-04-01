package commands

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
	"strings"
	"time"
)

func commandCatch(cfg *Config, args ...string) error {
	if len(args) < 1 {
		return errors.New("enter a {number/location} to explore")
	}

	pokemon := args[0]

	if !slices.Contains(cfg.ExploreList, pokemon) {
		return errors.New("invalid Pokemon. It is not in your explore list")
	}

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

	cfg.Theme.Success.Printf("%s %s\n", cfg.ThemeFunc.Pokemon(pokemon), cfg.ThemeFunc.Success("was caught!!"))
	cfg.Theme.Prompt.Printf("What would you like to name your %s? ", pokemon)

	var nickname string
	if cfg.Scanner.Scan() {
		nickname = cfg.Scanner.Text()
	}

	if nickname == "" {
		nickname = numbered(cfg, pokemon)
		cfg.Theme.Success.Printf("Got it! Sent %s to the Pokedex.\n", nickname)
	} else {
		cfg.Theme.Success.Printf("Got it! Sent %s to the Pokedex.\n", nickname)
	}
	pokemonStruct.NickName = nickname

	cfg.Pokedex = append(cfg.Pokedex, pokemonStruct)

	return nil
}

func numbered(cfg *Config, pokemonName string) string {
	total := 0
	for _, pokemon := range cfg.Pokedex {
		if strings.EqualFold(pokemon.NickName, pokemonName) {
			total++
		}
	}
	if total == 0 {
		return pokemonName
	}
	total++
	return fmt.Sprintf("%s %d", pokemonName, total)
}
