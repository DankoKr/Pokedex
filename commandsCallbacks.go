package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	availableCommands := getCommands()
	for _, cmd := range availableCommands{
        fmt.Printf("- %v: %v\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(cfg *config, args ...string) error{
	fmt.Println("Exiting Pokedex....")
    os.Exit(0)
	return nil
}

func commandMap(cfg *config, args ...string) error{
	resp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	fmt.Printf("Location areas: \n")
	for _, area := range resp.Results {
        fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationsURL = resp.Next
	cfg.previousLocationsURL = resp.Previous

    return nil
}

func commandMapB(cfg *config, args ...string) error{
    if cfg.previousLocationsURL == nil {
		return errors.New("this is the first page")
	}

	resp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)
	if err != nil {
		return err
	}

	fmt.Printf("Previous Location areas: \n")
	for _, area := range resp.Results {
        fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationsURL = resp.Next
	cfg.previousLocationsURL = resp.Previous

    return nil
}

func commandExplore(cfg *config, args ...string) error {
    if len(args) != 1 {
		return errors.New("no location provided")
	}

	locationName := args[0]

   locationArea, err := cfg.pokeapiClient.GetLocationArea(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemons in %s: \n", locationArea.Name)
	for _, pokemon := range locationArea.PokemonEncounters {
        fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
    return nil
}

func commandCatch(cfg *config, args ...string) error {
    if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

   pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s\n", pokemonName)

	treshHold := 50
    randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > treshHold {
		return fmt.Errorf("%s was not caught", pokemonName)
	}

    fmt.Printf("%s was caught\n", pokemonName)
    cfg.catchedPokemons[pokemonName] = pokemon

    return nil
}

func commandInspect(cfg *config, args ...string) error {
    if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

   pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	pokemon, ok := cfg.catchedPokemons[pokemonName] 
	if !ok {
		return errors.New("pokemon was not found in the Pokedex")
	}

	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats: ")
	for _, stat := range pokemon.Stats{
		fmt.Printf("- %s: %v\n", stat.Stat.Name, stat.BaseStat)
	} 

	fmt.Println("Types: ")
	for _, stat := range pokemon.Types{
		fmt.Printf("- %s\n", stat.Type.Name)
	} 

    return nil
}