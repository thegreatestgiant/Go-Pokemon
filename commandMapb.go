package main

import (
	"errors"
	"fmt"
)

func commandMapb(cfg *config) error {
	if cfg.previous == nil {
		return errors.New("you are on the first page so there are no previous locations")
	}

	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.previous)
	if err != nil {
		return err
	}

	cfg.next = resp.Next
	cfg.previous = resp.Previous

	fmt.Println("Areas")
	for _, area := range resp.Results {
		fmt.Printf("  - %s\n", area.Name)
	}

	return nil
}
