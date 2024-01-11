package main

import "github.com/DankoKr/Pokedex/internal/pokeapi"

type config struct {
	pokeapiClient  pokeapi.Client
	nextLocationsURL *string
	previousLocationsURL *string
}

func main() {
	cfg := config {
       pokeapiClient: pokeapi.NewClient(),
	}
	startRepl(&cfg)
}