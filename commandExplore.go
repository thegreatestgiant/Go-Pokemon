package main

import (
	"errors"
	"fmt"
	"strconv"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("enter a {number/location} to explore")
	}

	area := args[0]
	if len(args[0]) < 3 {
		resp, err, _ := cfg.pokeapiClient.GetLocationAreas(cfg.current)
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

	areaResp, err := cfg.pokeapiClient.GetAreaResp(&area)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", areaResp.Name)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range areaResp.PokemonEncounters {
		fmt.Printf("  - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
