package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() 
}


func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
    "help": {
        name:        "help",
        description: "Displays a help message",
        callback:    commandHelp,
    },
    "exit": {
        name:        "exit",
        description: "Exits the Pokedex",
        callback:    commandExit,
    },
}
}


func startRepl() {
    scanner := bufio.NewScanner(os.Stdin)
	availableCommands := getCommands()

	for {
        fmt.Print("pokedex > ")
	    scanner.Scan()
	    inputCommand := scanner.Text()

        command, ok := availableCommands[inputCommand]
		if !ok {
			fmt.Println("Command not recognised")
			continue
		}

		command.callback()
	}

}