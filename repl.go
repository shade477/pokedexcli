package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	for ;; {
		fmt.Print("pokedex> ")

		s := bufio.NewScanner(os.Stdin)
		s.Scan()

		if s.Err() != nil {
			fmt.Print(s.Err().Error())
		}

		//Parses input and prints first word
		// words := cleanInput(s.Text())[0] 
		// if len(words) == 0 {
		// 	continue
		// }

		// fmt.Printf("Your command was: %s\n", words)

		commandExec(s.Text())
	}
}

func commandExit() error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func getHelp() error {
	fmt.Println(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}

func cleanInput(text string) []string {
	res := strings.Split(strings.Trim(strings.ToLower(text), " "), " ")
	return res
}