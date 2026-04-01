package main

import (
	"time"

	"github.com/thegreatestgiant/Go-Pokemon/internal/commands"
	"github.com/thegreatestgiant/Go-Pokemon/internal/pokeapi"
	"github.com/thegreatestgiant/Go-Pokemon/internal/theme"
)

func main() {
	appTheme := theme.LoadTheme()
	appThemeFunc := theme.LoadThemeFunc()

	cfg := commands.Config{
		PokeapiClient: pokeapi.NewClient(time.Hour),
		Pokedex:       make(map[string]pokeapi.Pokemon),
		Theme:         appTheme,
		ThemeFunc:     appThemeFunc,
		Debug:         pokeapi.NewClient(time.Hour).Debug,
		Art:           true,
	}

	startRepl(&cfg)
}
