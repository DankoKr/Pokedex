package main

import (
	"fmt"
	"os"
)

func commandHelp() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	availableCommands := getCommands()
	for _, cmd := range availableCommands{
        fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
}

func commandExit() {
	fmt.Println("Exiting Pokedex....")
    os.Exit(0)
}