package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/oakinh/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for scanner.Scan() {

		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}
		allCommands := getCommands()
		command, ok := allCommands[text[0]]
		if !ok {
			fmt.Println("Invalid command. Type 'help' for a list of commands")
		} else {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
		fmt.Print("Pokedex > ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Reading input:", err)
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas of the game map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas of the game map",
			callback:    commandMapb,
		},
	}
}

func cleanInput(input string) []string {
	output := strings.ToLower(input)
	words := strings.Fields(output)
	return words
}
