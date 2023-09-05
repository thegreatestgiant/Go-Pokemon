package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	// Add caching
	resp, current, err := cfg.pokeapiClient.GetLocationAreas(cfg.next)
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

	resp, current, err := cfg.pokeapiClient.GetLocationAreas(cfg.previous)
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
