package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := getConfig()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if scanner.Err() != nil {
			break
		}
		input := scanner.Text()
		firstWord := cleanInput(input)[0]
		secondWord := ""
		if len(cleanInput(input)) > 1 {
			secondWord = cleanInput(input)[1]
		}
		commands := config.commands
		command, ok := commands[firstWord]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(&config, secondWord)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

	}
}
