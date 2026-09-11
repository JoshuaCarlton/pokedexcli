package main

import (
	"fmt"
)

func commandInspect(config *config, pokemon string) error {
	pokeData, ok := config.caught[pokemon]
	if !ok {
		return fmt.Errorf("you have not caught %s", pokemon)
	}
	fmt.Printf("Name: %s\n", pokeData.Name)
	fmt.Printf("Height: %d\n", pokeData.Height)
	fmt.Printf("Weight: %d\n", pokeData.Weight)
	fmt.Printf("Stats: \n")
	for _, stat := range pokeData.Stats {
		fmt.Printf("	-%s: %d\n", stat.StatInfo.Name, stat.BaseStat)
	}
	fmt.Printf("Types: \n")
	for _, pokeType := range pokeData.Types {
		fmt.Printf("	-%s\n", pokeType.TypeInfo.Name)
	}
	return nil
}
