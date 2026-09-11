package main

import (
	"encoding/json"
	"fmt"
)

type exoloreData struct {
	Encounters []encounter `json:"pokemon_encounters"`
}

type encounter struct {
	Pokemon pokemon `json:"pokemon"`
}

type pokemon struct {
	Name string `json:"name"`
}

func commandExplore(config *config, area string) error {
	if area == "" {
		return fmt.Errorf("must provide an area to explore")
	}
	url := "https://pokeapi.co/api/v2/location-area/"
	full_url := url + area
	bytes, err := get_bytes(config, full_url)
	if err != nil {
		return err
	}

	var data exoloreData

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")
	for _, encounter := range data.Encounters {
		fmt.Println("-", encounter.Pokemon.Name)
	}
	return nil
}
