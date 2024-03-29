package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}


func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
    "help": {
        name:        "help",
        description: "Displays a help message",
        callback:    commandHelp,
    },
	"map": {
        name:        "map",
        description: "Lists locations",
        callback:    commandMap,
    },
	"mapb": {
        name:        "mapb",
        description: "Lists previous locations",
        callback:    commandMapB,
    },
	"explore": {
        name:        "explore {location_name}",
        description: "Explore a location",
        callback:    commandExplore,
    },
	"catch": {
        name:        "catch {pokemon_name}",
        description: "Catch a pokemon",
        callback:    commandCatch,
    },
	"inspect": {
        name:        "inspect {pokemon_name}",
        description: "Inspect a pokemon stats",
        callback:    commandInspect,
    },
	"pokedex": {
        name:        "pokedex",
        description: "Displays all the caught Pokemons",
        callback:    commandPokedex,
    },
    "exit": {
        name:        "exit",
        description: "Exits the Pokedex",
        callback:    commandExit,
    },
}
}


func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}