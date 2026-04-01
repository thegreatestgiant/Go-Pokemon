package state

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/thegreatestgiant/Go-Pokemon/internal/commands"
	"github.com/thegreatestgiant/Go-Pokemon/internal/theme"
)

type Settings struct {
	DebugMode    bool         `json:"debug_mode"`
	ShowAsciiArt bool         `json:"pokemon_as_ascii"`
	ThemeColors  theme.Colors `json:"theme_colors"`
}

func SaveSettings(cfg *commands.Config) bool {
	settings := Settings{
		DebugMode:    cfg.Debug,
		ShowAsciiArt: cfg.Art,
		ThemeColors:  theme.DefaultColors,
	}

	jsonFile, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		if settings.DebugMode {
			cfg.Theme.Error.Printf("Failed to Marshal Settings: %v\n", err)
		}
		return false
	}

	err = os.WriteFile("settings.json", jsonFile, 0o644)
	if err != nil {
		if settings.DebugMode {
			cfg.Theme.Error.Printf("Couldn't write to file: %v\n", err)
		}
		return false
	}
	return true
}

func loadSettings() *Settings {
	settings := Settings{
		DebugMode:    false,
		ShowAsciiArt: true,
	}

	file, err := os.ReadFile("settings.json")
	if err != nil {
		if settings.DebugMode {
			// cfg.Theme.Error.Printf("Couldn't write to file: %v\n", err)
			fmt.Printf("Couldn't read file: %v\n", err)
		}
		return &settings
	}

	err = json.Unmarshal(file, &settings)
	if err != nil {
		if settings.DebugMode {
			// cfg.Theme.Error.Printf("Couldn't write to file: %v\n", err)
			fmt.Printf("Couldn't marshal json file: %v\n", err)
		}
		return &settings
	}
	return &settings
}
