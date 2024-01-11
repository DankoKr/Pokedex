package main

import (
	"errors"
	"fmt"
	"os"
)

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	availableCommands := getCommands()
	for _, cmd := range availableCommands{
        fmt.Printf("- %v: %v\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(cfg *config) error{
	fmt.Println("Exiting Pokedex....")
    os.Exit(0)
	return nil
}

func commandMap(cfg *config) error{
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

func commandMapB(cfg *config) error{
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