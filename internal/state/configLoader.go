package state

import (
	"time"

	"github.com/thegreatestgiant/Go-Pokemon/internal/commands"
	"github.com/thegreatestgiant/Go-Pokemon/internal/pokeapi"
	"github.com/thegreatestgiant/Go-Pokemon/internal/theme"
)

func GetConfig() commands.Config {
	settings := loadSettings()

	appTheme := theme.LoadTheme(&settings.ThemeColors)
	appThemeFunc := theme.LoadThemeFunc(&settings.ThemeColors)

	debugging := settings.DebugMode
	showArt := settings.ShowAsciiArt

	pokedex := loadPokedex(debugging)

	return commands.Config{
		PokeapiClient: pokeapi.NewClient(time.Hour, debugging),
		Pokedex:       pokedex,
		Theme:         appTheme,
		ThemeFunc:     appThemeFunc,
		Debug:         debugging,
		Art:           showArt,
	}
}
