package main

import "github.com/thegreatestgiant/Go-Pokemon/internal/state"

func main() {
	cfg := state.GetConfig()
	startRepl(&cfg)
}
