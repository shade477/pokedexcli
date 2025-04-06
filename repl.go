package main

import "strings"

func cleanInput(text string) []string {
	res := strings.Split(strings.Trim(strings.ToLower(text), " "), " ")
	return res
}