package state

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/thegreatestgiant/Go-Pokemon/internal/commands"
	"github.com/thegreatestgiant/Go-Pokemon/internal/pokeapi"
)

type SaveGame struct {
	Pokedex []pokeapi.Pokemon
}

func loadPokedex(isDebug bool) []pokeapi.Pokemon {
	file, err := os.Open("pokedex.sav")
	if err != nil {
		if isDebug {
			fmt.Printf("err: %v\n", err)
		}
		return []pokeapi.Pokemon{}
	}

	var saveState SaveGame

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&saveState)
	if err != nil {
		if isDebug {
			err.Error()
		}
		return []pokeapi.Pokemon{}
	}

	return saveState.Pokedex
}

func SavePokedex(cfg *commands.Config) bool {
	file, err := os.Create("pokedex.sav")
	if err != nil {
		if cfg.Debug {
			cfg.Theme.Error.Printf("Gob Error: %v\n", err)
		}
		return false
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)

	err = encoder.Encode(&SaveGame{Pokedex: cfg.Pokedex})
	if err != nil {
		if cfg.Debug {
			cfg.Theme.Error.Printf("Gob Error: %v\n", err)
		}
		return false
	}

	return true
}
