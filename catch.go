package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type pokeData struct {
	Name   string     `json:"name"`
	BaseXp int        `json:"base_experience"`
	Height int        `json:"height"`
	Weight int        `json:"weight"`
	Stats  []stat     `json:"stats"`
	Types  []pokeType `json:"types"`
}

type pokeType struct {
	TypeInfo typeInfo `json:"type"`
}

type typeInfo struct {
	Name string `json:"name"`
}

type stat struct {
	BaseStat int      `json:"base_stat"`
	StatInfo statInfo `json:"stat"`
}

type statInfo struct {
	Name string `json:"name"`
}

func commandCatch(config *config, pokemon string) error {
	if pokemon == "" {
		return fmt.Errorf("need name of pokemon to catch")
	}
	url := "https://pokeapi.co/api/v2/pokemon/"
	full_url := url + pokemon
	bytes, err := get_bytes(config, full_url)
	if err != nil {
		return err
	}

	var pd pokeData

	err = json.Unmarshal(bytes, &pd)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pd.Name)
	randomFloat := rand.Float32()
	chanceToBeat := float32(pd.BaseXp) / 700.0
	if randomFloat >= float32(chanceToBeat) {
		fmt.Printf("%s was caught!\n", pd.Name)
		config.caught[pd.Name] = pd
	} else {
		fmt.Printf("%s escaped!\n", pd.Name)
	}
	return nil
}
