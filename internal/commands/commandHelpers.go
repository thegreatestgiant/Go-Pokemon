package commands

import (
	"strings"

	"github.com/thegreatestgiant/Go-Pokemon/internal/pokeapi"
)

func parseFlags(args ...string) (commandArgs []string, flag string, flagArgs []string) {
	tripped := false
	for _, arg := range args {
		if tripped {
			flagArgs = append(flagArgs, arg)
			continue
		}
		if strings.HasPrefix(arg, "-") {
			flag = arg[1:]
			for strings.HasPrefix(flag, "-") {
				flag = flag[1:]
			}
			tripped = true
		} else {
			commandArgs = append(commandArgs, arg)
		}
	}
	return commandArgs, flag, flagArgs
}

func (cfg *Config) findPokemon(nameSplice []string) (pokeapi.Pokemon, int, bool) {
	name := strings.Join(nameSplice, " ")
	p, i, f := pokeapi.Pokemon{}, -1, false
	for idx, pokemon := range cfg.Pokedex {
		if pokemon.NickName == name {
			return pokemon, idx, true
		} else if pokemon.Name == name {
			p = pokemon
			i = idx
			f = true
		}
	}
	return p, i, f
}
