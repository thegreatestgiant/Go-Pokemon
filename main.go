package main

import (
	"time"

	"github.com/thegreatestgiant/Go-Pokemon/internal/pokeapi"
	"github.com/thegreatestgiant/Go-Pokemon/internal/theme"
)

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	current       *string
	previous      *string
	pokedex       map[string]pokeapi.Pokemon
	theme         *theme.CLITheme
	themeFunc     *theme.CLIThemeFunc
	debug         bool
}

func main() {
	appTheme := theme.LoadTheme()
	appThemeFunc := theme.LoadThemeFunc()

	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		pokedex:       make(map[string]pokeapi.Pokemon),
		theme:         appTheme,
		themeFunc:     appThemeFunc,
		debug:         pokeapi.NewClient(time.Hour).Debug,
	}

	startRepl(&cfg)
}
