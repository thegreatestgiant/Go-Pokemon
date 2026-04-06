# 🎮 Go-Pokémon

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Platform](https://img.shields.io/badge/Platform-Terminal-black?style=for-the-badge&logo=gnubash&logoColor=white)
![API](https://img.shields.io/badge/API-PokéAPI-EF5350?style=for-the-badge)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

**A fully interactive Pokémon adventure — built in Go, played in your terminal.**

```sh
Pokemon > explore pastoria-city-area
  Exploring pastoria-city-area...
  Found Pokemon:
    - tentacool
    - tentacruel
    - magikarp

Pokemon > catch magikarp
  Throwing a Pokeball at magikarp...
  magikarp was caught!!
  What would you like to name your magikarp? Splash
  Got it! Sent Splash to the Pokedex.
```

</div>

---

## 📋 Table of Contents

- [Features](#-features)
- [Quick Start](#️-quick-start)
- [Gameplay](#️-gameplay)
  - [Commands](#commands)
  - [Catch Mechanic](#catch-mechanic)
- [Under the Hood](#️-under-the-hood)
  - [Go Patterns Used](#go-patterns-used)
- [Roadmap](#️-roadmap)
- [License](#️-license)

## ✨ Features

- 🌍 **World Navigation** — Page through Pokémon location areas, forward and backward
- 🎲 **Probability-Based Catching** — Catch difficulty scales with a Pokémon's real `BaseExperience` stat
- 📖 **Persistent Pokédex** — Your caught Pokémon survive between sessions via a binary save file
- 🔬 **Deep Inspection** — View stats, types, and full move lists for any Pokémon you've caught
- ⚡ **In-Memory Caching** — API responses are cached with TTL expiration, eliminating redundant network calls
- 🎨 **Theming** — Fully colorized terminal output with a configurable RGB theme engine
- 🖼️ **ASCII Art** — Pokémon sprites rendered as ASCII directly in your terminal

---

## ⚙️ Quick Start

**Prerequisites:** [Go 1.21+](https://go.dev/dl/) and an internet connection.

```bash
git clone https://github.com/thegreatestgiant/Go-Pokemon.git
cd Go-Pokemon
go run .
```

```bash
# Or build a binary
go build -o go-pokemon .
./go-pokemon

# Run tests
go test ./...
```

---

## 🕹️ Gameplay

### Commands

| Command | Argument | Description |
|---|---|---|
| `help` | — | List all commands |
| `map` | — | Show the next 20 location areas |
| `mapb` | — | Show the previous 20 location areas |
| `explore` | `{name}` or `{#}` | List all Pokémon in an area |
| `catch` | `{pokemon}` | Throw a Pokéball — success is not guaranteed |
| `release` | `{pokemon}` or `*` | Remove one or all Pokémon from your Pokédex |
| `inspect` | `{pokemon} [-m {X}]` | View stats; `-m` flag shows the first `X` moves |
| `pokedex` | — | List all caught Pokémon |
| `exit` | — | Save and quit |

### Catch Mechanic

Difficulty is determined by a Pokémon's real `BaseExperience` stat. A random number is drawn from `[0, BaseExperience)` — if it falls below `50`, the catch fails.

| Pokémon | BaseExperience | Approx. Catch Rate |
|---|---|---|
| Magikarp | 40 | ~100% |
| Gastly | 62 | ~81% |
| Dragonite | 270 | ~19% |
| Mewtwo | 340 | ~15% |

> Rare Pokémon are genuinely harder to catch — just like the mainline games.

---

## 🏗️ Under the Hood

```
┌──────────────────────────────────────────────┐
│                   main.go                    │
│        Bootstraps config, starts REPL        │
└──────────────────────┬───────────────────────┘
                       │
┌──────────────────────▼───────────────────────┐
│                   repl.go                    │
│     Read → Parse → Dispatch (command reg.)   │
└──────┬────────────────────────────┬──────────┘
       │                            │
┌──────▼──────┐             ┌───────▼──────────┐
│ command*.go │    ...      │   command*.go    │
│ catch,      │             │ map, explore,    │
│ inspect,    │             │ pokedex, release │
│ release     │             │                  │
└──────┬──────┘             └───────┬──────────┘
       │                            │
┌──────▼────────────────────────────▼──────────┐
│             internal/pokeapi/                │
│    HTTP Client · Request Methods · Structs   │
└──────────────────────┬───────────────────────┘
                       │
┌──────────────────────▼───────────────────────┐
│             internal/pokecache/              │
│     TTL Cache · Mutex Safety · Reaper        │
└──────────────────────────────────────────────┘
```

### Go Patterns Used

| Pattern | Where | Purpose |
|---|---|---|
| **First-class functions** | `cliCommand.Callback` | Uniform command dispatch without a switch statement |
| **Goroutines + `sync.Mutex`** | `pokecache` reaper loop | Background TTL expiration without blocking the REPL |
| **`time.Ticker`** | Cache reaper | Efficient, non-blocking periodic execution |
| **Pointer receivers** | All `Client` and `Cache` methods | Mutates shared state without copying |
| **`encoding/json` + `encoding/gob`** | API unmarshaling, save file | Type-safe API consumption; compact binary persistence |
| **Variadic callbacks** | `args ...string` on all commands | Clean, uniform dispatch signature across all commands |
| **`bufio.Scanner`** | REPL input | Handles arbitrary-length input and mid-catch nickname prompts |
| **Struct composition** | `Config` embeds client, cache, theme | Avoids inheritance; Go-idiomatic dependency passing |

Powered by [PokéAPI v2](https://pokeapi.co/) — free, open, no key required.

---

## 🗺️ Roadmap

- [x] Persistent save data (Gob-encoded binary)
- [x] Pokémon release (`release *` clears all)
- [x] `inspect --moves` flag with configurable limit
- [x] Colored terminal output via RGB theme engine
- [x] ASCII sprite rendering on inspect
- [ ] Battle system — turn-based combat using existing move/stat data
- [ ] Shiny Pokémon — additional RNG layer on successful catches
- [ ] Pokémon evolution — trigger evolutions based on in-game conditions
- [ ] Interactive TUI — upgrade REPL to full terminal UI with [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [ ] Settings toggles as REPL commands — expose `debug` and `ascii-art` without editing JSON
- [ ] Configurable cache TTL — via CLI flag or settings file (currently hardcoded to 1 hour)
- [ ] Expanded test coverage — `pokecache_test.go` has scaffolding; real table-driven tests would complete the suite

---

## 📄 License

MIT — see [`LICENSE`](LICENSE) for details.

---

<div align="center">
Built with ❤️ and Go · Powered by <a href="https://pokeapi.co/">PokéAPI</a>
</div>
