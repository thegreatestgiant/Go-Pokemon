package main

import (
	"errors"
	"fmt"

	"github.com/thegreatestgiant/Go-Pokemon/internal/pokeapi"
)

func commandMapb(cfg *config) error {
	pokeApiClient := pokeapi.NewClient()

	if cfg.previous == nil {
		return errors.New("you are on the first page so there are no previous locations")
	}

	resp, err := pokeApiClient.GetLocationAreas(cfg.previous)
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
