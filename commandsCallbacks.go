package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DankoKr/Pokedex/internal/pokeapi"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	availableCommands := getCommands()
	for _, cmd := range availableCommands{
        fmt.Printf("- %v: %v\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit() error{
	fmt.Println("Exiting Pokedex....")
    os.Exit(0)
	return nil
}

func commandMap() error{
	pokeapiClient := pokeapi.NewClient()
	resp, err := pokeapiClient.ListLocations()
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("Location areas: ")
	for _, area := range resp.Results {
        fmt.Printf("- %s\n", area.Name)
	}

    return nil
}