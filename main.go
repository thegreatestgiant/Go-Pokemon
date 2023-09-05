package main

import (
	"time"

	"github.com/thegreatestgiant/Go-Pokemon/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	current       *string
	previous      *string
	pokedex       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		pokedex:       make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
