# 🎮 Go-Pokémon

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Platform](https://img.shields.io/badge/Platform-Terminal-black?style=for-the-badge&logo=gnubash&logoColor=white)
![API](https://img.shields.io/badge/API-PokéAPI-EF5350?style=for-the-badge)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

**A fully interactive Pokémon adventure game — built entirely in Go, played entirely in your terminal.**

Explore the Pokémon world, battle the odds to catch wild Pokémon, and build your ultimate Pokédex — all powered by a live REST API, an in-memory caching layer, and a hand-rolled REPL engine.

[Getting Started](#%EF%B8%8F-installation--setup) · [How to Play](#%EF%B8%8F-how-to-play) · [Architecture](#%EF%B8%8F-architecture) · [Tech Stack](#%EF%B8%8F-tech-stack--go-patterns-used)

</div>

---

## 🗺️ Gameplay Preview

```
Pokemon > map
  Cache Missed
  Areas
    1. canalave-city-area
    2. eterna-city-area
    3. pastoria-city-area
    ...

Pokemon > explore 3
  Exploring pastoria-city-area...
  Found Pokemon:
    - tentacool
    - tentacruel
    - magikarp

Pokemon > catch magikarp
  Cache Missed
  Throwing a Pokeball at magikarp...
  magikarp was caught!!

Pokemon > inspect magikarp
  Name: magikarp
  Height: 9
  Weight: 100
  Stats:
    - hp: 20
    - attack: 10
    - defense: 55
  Types:
    - water

Pokemon > pokedex
  Your Pokedex:
    - magikarp
```

---

## ✨ Features

- 🌍 **World Navigation** — Page through 20 Pokémon location areas at a time, forward and backward
- 🔍 **Area Exploration** — Inspect any location to see which wild Pokémon can be encountered there
- 🎲 **Catch Mechanic with Real RNG** — Catch attempts use the Pokémon's actual `BaseExperience` stat to calculate difficulty; rare Pokémon are genuinely harder to catch
- 📖 **Live Pokédex** — Your caught Pokémon are stored in-session with full stat data
- 🔬 **Pokémon Inspector** — View height, weight, base stats, and types for any Pokémon in your Pokédex
- ⚡ **Smart API Caching** — Identical API calls are served from an in-memory cache instead of hitting the network — includes background TTL expiration
- 🔤 **Case-Insensitive Input** — Type commands in any capitalization
- 🔢 **Flexible Explore Input** — Explore areas by number *or* by name

---

## 🕹️ How to Play

### Starting the Game

Once running (see [Installation](#%EF%B8%8F-installation--setup)), you'll be greeted with the prompt:

```
Pokemon >
```

Type any command below and press **Enter**.

---

### Command Reference

| Command | Argument | Description |
|---|---|---|
| `help` | — | Displays all available commands and their descriptions |
| `map` | — | Lists the **next** 20 Pokémon location areas in the world |
| `mapb` | — | Lists the **previous** 20 Pokémon location areas |
| `explore` | `{name}` or `{number}` | Lists all Pokémon found in a specific area |
| `catch` | `{pokemon-name}` | Throws a Pokéball at the named Pokémon |
| `inspect` | `{pokemon-name}` | Shows stats for a Pokémon **already in your Pokédex** |
| `pokedex` | — | Lists the names of all Pokémon you have caught |
| `exit` | — | Exits the game |

---

### Step-by-Step Walkthrough

**1. Discover the world with `map`**

```
Pokemon > map
```

This fetches the first 20 location areas from the Pokémon world. Run it again to see the next 20. Use `mapb` to go back.

**2. Explore an area using its number or name**

```
Pokemon > explore 5
```

Uses the 5th result from your last `map` call. Alternatively, use the full area name:

```
Pokemon > explore pastoria-city-area
```

**3. Catch a Pokémon**

```
Pokemon > catch gastly
```

The game fetches `gastly`'s data from PokéAPI and calculates a catch probability based on its `BaseExperience` value. Pokémon with high `BaseExperience` (like legendaries) are much harder to catch — just like in the real games.

> 💡 **Catch Mechanic:** A random number is generated between `0` and the Pokémon's `BaseExperience`. If that number exceeds `50`, the catch fails. A Pokémon with `BaseExperience = 64` (like Gastly) gives you roughly a 78% chance. A Pokémon with `BaseExperience = 340` (like Mewtwo) gives you only about a 15% chance.

**4. Inspect a caught Pokémon**

```
Pokemon > inspect gastly
```

Displays the Pokémon's name, height, weight, base stats (HP, Attack, Defense, etc.), and types. Only works for Pokémon already in your Pokédex.

**5. View your full Pokédex**

```
Pokemon > pokedex
```

Lists every Pokémon you've successfully caught in this session.

---

### Tips & Rules

- 🔄 **Sessions are not persistent** — your Pokédex resets when you exit the game.
- ✏️ **Input is case-insensitive** — `Pikachu`, `PIKACHU`, and `pikachu` are all valid.
- ⚠️ **You must `explore` before you can `catch`** — you need to know what's in an area first!
- 🔢 **The number shortcut for `explore`** works based on the *most recently viewed* `map` page.
- ♻️ **Re-running the same `map` or `explore`** command? It'll be served from cache — near-instant response.

---

## ⚙️ Installation & Setup

### Prerequisites

- [Go 1.21+](https://go.dev/dl/) installed on your machine
- A working internet connection (for the initial API calls; subsequent calls use the cache)

### Clone and Run

```bash
# Clone the repository
git clone https://github.com/thegreatestgiant/Go-Pokemon.git
cd Go-Pokemon

# Run directly
go run .

# OR build a binary first
go build -o go-pokemon .
./go-pokemon
```

### Run Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test ./... -v
```

---

## 🏗️ Architecture

The project is organized into a clean, layered architecture where each layer has a single responsibility.

```
┌─────────────────────────────────────────────────────────┐
│                        main.go                          │
│         Initializes config state, starts REPL           │
└───────────────────────┬─────────────────────────────────┘
                        │
┌───────────────────────▼─────────────────────────────────┐
│                       repl.go                           │
│   Read → Parse → Dispatch loop (command registration)   │
└───────┬───────────────────────────────────┬─────────────┘
        │                                   │
┌───────▼───────┐                   ┌───────▼───────────────┐
│ command*.go   │                   │   command*.go         │
│  (catch,      │       ...         │  (map, explore,       │
│   inspect)    │                   │   pokedex, help)      │
└───────┬───────┘                   └──────────┬────────────┘
        │                                      │
┌───────▼──────────────────────────────────────▼────────────┐
│                   internal/pokeapi/                        │
│      HTTP Client  •  Request Methods  •  Type Structs      │
└───────────────────────────┬────────────────────────────────┘
                            │
┌───────────────────────────▼────────────────────────────────┐
│                  internal/pokecache/                        │
│       In-memory TTL Cache  •  Mutex Safety  •  Reaper      │
└─────────────────────────────────────────────────────────────┘
```

### Key Files

| File | Role |
|---|---|
| `main.go` | Entry point; defines `config` (global app state); sets 1-hour cache TTL |
| `repl.go` | REPL engine; input parsing; command registry (`getCommands`) |
| `commandCatch.go` | RNG-based catch logic using `BaseExperience` |
| `commandExplore.go` | Dual-mode area lookup (number index or name string) |
| `commandMap.go` | Paginated location browsing, forward and backward |
| `commandInspect.go` | Pokédex stat viewer |
| `internal/pokeapi/pokeapi.go` | HTTP client constructor with configurable timeout |
| `internal/pokeapi/pokemon_req.go` | Fetches full Pokémon data with cache integration |
| `internal/pokeapi/location_areas_req.go` | Paginated location area fetching |
| `internal/pokeapi/explore_req.go` | Single location area detail fetching |
| `internal/pokecache/pokecache.go` | TTL cache with mutex protection and background reaper goroutine |

---

## 🛠️ Tech Stack & Go Patterns Used

This project was built using only the **Go standard library** — no third-party frameworks.

### Go Features & Patterns

| Pattern / Feature | Where It's Used | Why It Matters |
|---|---|---|
| **Variadic functions** | All `command*` callbacks (`args ...string`) | Uniform command dispatch signature |
| **Pointer receivers & method sets** | `(c *Client)` on all API methods; `(c *Cache)` on cache methods | Mutating shared state without copying |
| **Interface-ready design** | `config` passes a concrete `pokeapi.Client` by value | Commands are decoupled from HTTP implementation |
| **Goroutines** | `go c.reapLoop(dur)` in `pokecache.NewCache` | Background cache expiration without blocking |
| **`sync.Mutex`** | Guards all `cache` map reads/writes | Prevents data races in concurrent access |
| **`time.Ticker`** | Drives the cache reaper loop | Efficient, non-blocking periodic execution |
| **`encoding/json`** | Unmarshals all API responses into typed structs | Type-safe API consumption |
| **First-class functions** | `callback func(*config, ...string) error` in `cliCommand` | Enables the command registry pattern |
| **`bufio.Scanner`** | REPL input reading | Handles arbitrary-length input safely |
| **`strings.Fields`** | Input tokenization | Whitespace-agnostic parsing |
| **`math/rand`** | Catch mechanic RNG | Gameplay probability scaling |
| **Struct embedding / composition** | `config` composes `pokeapi.Client` and `pokecache.Cache` | Avoids inheritance in favor of composition |

### External API

| Resource | Usage |
|---|---|
| [PokéAPI v2](https://pokeapi.co/) | All Pokémon and location data — fully free, no API key required |

---

## 🗺️ Roadmap

Potential features for future development:

- [X] **Persistent save data** — Write the Pokédex to a JSON file between sessions
- [ ] **Battle system** — Turn-based combat using actual move and stat data (already in the `Pokemon` struct)
- [ ] **Shiny Pokémon** — Additional rare RNG layer on successful catches
- [ ] **Pokémon release** — Remove a Pokémon from your Pokédex
- [ ] **Pokémon evolution** — Trigger evolutions based on in-game conditions
- [ ] **`inspect --moves` flag** — Display the full move list from the already-fetched data
- [x] **Colored terminal output** — Use ANSI codes or a lightweight package for a richer UX
- [ ] **Interactive TUI** — Upgrade the REPL to a full terminal UI with [Bubble Tea](https://github.com/charmbracelet/bubbletea)

---

## 📄 License

This project is licensed under the MIT License. See [`LICENSE`](LICENSE) for details.

---

<div align="center">

Built with ❤️ and Go · Powered by [PokéAPI](https://pokeapi.co/)

</div>
