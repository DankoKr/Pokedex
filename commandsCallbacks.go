package main

import (
	"errors"
	"fmt"
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