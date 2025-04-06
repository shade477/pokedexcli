package main

import "fmt"

type cliCommand struct {
	name		string
	description string
	callback	func() error
}

func commandExec(command string) {
	commands := map[string]cliCommand{
		"exit": {
			name:			"exit",
			description:	"Exit the pokedex",
			callback:		commandExit,
		},
		"help": {
			name:			"help",
			description: 	"Displays a help message",
			callback: 		getHelp,
		},
	}

	c, exists:= commands[command]
	if !exists{
		fmt.Println("Unknown command")
		return
	}
	c.callback()
}