package commands

import (
	"errors"
	"fmt"
)

func commandMap(cfg *Config, args ...string) error {
	// Add caching
	resp, current, err := cfg.PokeapiClient.GetLocationAreas(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Next = resp.Next
	cfg.Current = &current
	cfg.Previous = resp.Previous

	cfg.Theme.Header.Println("Areas")
	for i, area := range resp.Results {
		fmt.Printf("  %v. %s\n", i+1, cfg.ThemeFunc.Location(area.Name))
	}

	return nil
}

func commandMapb(cfg *Config, args ...string) error {
	if cfg.Previous == nil {
		return errors.New("you are on the first page. There are no previous locations")
	}

	resp, current, err := cfg.PokeapiClient.GetLocationAreas(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = resp.Next
	cfg.Current = &current
	cfg.Previous = resp.Previous

	cfg.Theme.Header.Println("Areas")
	for i, area := range resp.Results {
		fmt.Printf("  %v. %s\n", i+1, cfg.ThemeFunc.Location(area.Name))
	}

	return nil
}
