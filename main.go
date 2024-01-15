package main

import (
	"time"

	"github.com/DankoKr/Pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient  pokeapi.Client
	nextLocationsURL *string
	previousLocationsURL *string
	catchedPokemons map[string]pokeapi.Pokemon
}

func main() {
	cfg := config {
       pokeapiClient: pokeapi.NewClient(time.Hour),
	   catchedPokemons: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}