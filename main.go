package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Hello, World!")
}

func cleanInput(text string) []string {
	res := strings.Split(strings.Trim(text, " "), " ")
	return res
}