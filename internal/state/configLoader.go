package state

import (
	"time"

	"github.com/thegreatestgiant/Go-Pokemon/internal/commands"
	"github.com/thegreatestgiant/Go-Pokemon/internal/pokeapi"
	"github.com/thegreatestgiant/Go-Pokemon/internal/theme"
)

func GetConfig() commands.Config {
	appTheme := theme.LoadTheme()
	appThemeFunc := theme.LoadThemeFunc()

	pokedex := []pokeapi.Pokemon{}

	debugging := false
	showArt := true

	return commands.Config{
		PokeapiClient: pokeapi.NewClient(time.Hour, debugging),
		Pokedex:       pokedex,
		Theme:         appTheme,
		ThemeFunc:     appThemeFunc,
		Debug:         debugging,
		Art:           showArt,
	}
}
