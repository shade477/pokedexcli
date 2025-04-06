package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for ;; {
		fmt.Print("pokedex> ")
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		if s.Err() != nil {
			fmt.Print(s.Err().Error())
		}
		fmt.Printf("Your command was: %s\n", cleanInput(s.Text())[0])
	}
}
