package main

import "fmt"

func commandPokedex(config *config, _ string) error {
	fmt.Println("your pokedex:")
	for pokemon, _ := range config.caught {
		fmt.Printf("-%s\n", pokemon)
	}
	return nil
}
