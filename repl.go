package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
    "exit": {
        name:        "exit",
        description: "Exits the Pokedex",
        callback:    commandExit,
    },
}
}


func startRepl(cfg *config) {
    scanner := bufio.NewScanner(os.Stdin)
	availableCommands := getCommands()

	for {
        fmt.Print("pokedex > ")
	    scanner.Scan()
	    inputCommand := scanner.Text()

		if len(inputCommand) == 0 {
			continue
		}

        command, ok := availableCommands[inputCommand]
		if !ok {
			fmt.Println("Command not recognised")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}

}