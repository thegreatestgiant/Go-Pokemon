package commands

import (
	"errors"
	"fmt"
	"strconv"
)

func commandExplore(cfg *Config, args ...string) error {
	if len(args) < 1 {
		return errors.New("enter a {number/location} to explore")
	}

	area := args[0]
	if len(args[0]) < 3 {
		resp, _, err := cfg.PokeapiClient.GetLocationAreas(cfg.Current)
		if err != nil {
			return err
		}
		areaNum, err := strconv.Atoi(area)
		if err != nil {
			return err
		}
		areaNum -= 1
		area = resp.Results[areaNum].Name
	}

	areaResp, err := cfg.PokeapiClient.GetAreaResp(&area)
	if err != nil {
		return err
	}

	cfg.Theme.Header.Printf("Exploring %s...\n", cfg.ThemeFunc.Location(areaResp.Name))
	cfg.Theme.Info.Println("Found Pokemon:")
	cfg.ExploreList = nil
	for _, pokemon := range areaResp.PokemonEncounters {
		fmt.Printf("  - %s\n", cfg.ThemeFunc.Pokemon(pokemon.Pokemon.Name))
		cfg.ExploreList = append(cfg.ExploreList, pokemon.Pokemon.Name)
	}

	return nil
}
