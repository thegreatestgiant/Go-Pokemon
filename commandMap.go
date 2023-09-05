package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.next)
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
