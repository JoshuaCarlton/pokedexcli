package main

import "strings"

func cleanInput(text string) []string {
	words := strings.Fields(text)
	var lowers []string
	for _, word := range words {
		lowers = append(lowers, strings.ToLower(word))
	}
	return lowers
}
