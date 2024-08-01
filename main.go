package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	allCommands := getCommands()
	fmt.Printf(
		`Welcome to the Pokedex!
Usage:
	
%v: %v
%v: %v

`, allCommands["help"].name,
		allCommands["help"].description,
		allCommands["exit"].name,
		allCommands["exit"].description)
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
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
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for scanner.Scan() {

		text := scanner.Text()
		allCommands := getCommands()
		command, ok := allCommands[text]
		if !ok {
			fmt.Println("Invalid command. Type 'help' for a list of commands")
		} else {
			err := command.callback()
			if err != nil {
				fmt.Println("Error", err)
			}
		}
		fmt.Print("Pokedex > ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Reading input:", err)
	}
}
