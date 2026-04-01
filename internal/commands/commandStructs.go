package commands

import (
	"github.com/thegreatestgiant/Go-Pokemon/internal/pokeapi"
	"github.com/thegreatestgiant/Go-Pokemon/internal/theme"
)

type Config struct {
	PokeapiClient pokeapi.Client
	Next          *string
	Current       *string
	Previous      *string
	Pokedex       []pokeapi.Pokemon
	Theme         *theme.CLITheme
	ThemeFunc     *theme.CLIThemeFunc
	Debug         bool
	Art           bool
}

type cliCommand struct {
	name        string
	argument    string
	description string
	Callback    func(*Config, ...string) error
	priotity    int
}
