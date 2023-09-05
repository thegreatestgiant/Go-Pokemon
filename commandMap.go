package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func commandMap(cfg *config, args ...string) error {
	// Add caching
	resp, err, current := cfg.pokeapiClient.GetLocationAreas(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = resp.Next
	cfg.current = &current
	cfg.previous = resp.Previous

	fmt.Println("Areas")
	for i, area := range resp.Results {
		fmt.Printf("  %v. %s\n", i+1, area.Name)
	}

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previous == nil {
		return errors.New("you are on the first page so there are no previous locations")
	}

	resp, err, _ := cfg.pokeapiClient.GetLocationAreas(cfg.previous)
	if err != nil {
		return err
	}

	cfg.next = resp.Next
	cfg.previous = resp.Previous

	fmt.Println("Areas")
	for _, area := range resp.Results {
		lenOUrl := len(area.URL) - 4
		areaStr := area.URL[lenOUrl:]
		areaStr = strings.TrimPrefix(areaStr, "a")
		areaStr = strings.TrimPrefix(areaStr, "/")
		areaStr = strings.TrimSuffix(areaStr, "/")
		areaNum, err := strconv.Atoi(areaStr)
		if err != nil {
			return err
		}
		areaNum %= 20
		fmt.Printf("  %v. %s\n", areaNum, area.Name)
	}

	return nil
}
