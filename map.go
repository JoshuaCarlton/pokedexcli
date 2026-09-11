package main

import (
	"encoding/json"
	"fmt"
)

type mapData struct {
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []location `json:"results"`
}

type location struct {
	Name string `json:"name"`
}

func commandMapb(config *config, _ string) error {
	if config.mapb == "" {
		fmt.Println("you are on the first page")
		return nil
	}
	bytes, err := get_bytes(config, config.mapb)
	if err != nil {
		return err
	}

	var md mapData

	err = json.Unmarshal(bytes, &md)
	if err != nil {
		return err
	}
	config.mapn = md.Next
	config.mapb = md.Previous
	results := md.Results
	for _, result := range results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMap(config *config, _ string) error {
	if config.mapn == "" {
		fmt.Println("you are on the last page")
		return nil
	}
	bytes, err := get_bytes(config, config.mapn)
	if err != nil {
		return err
	}

	var md mapData

	err = json.Unmarshal(bytes, &md)
	if err != nil {
		return fmt.Errorf("couldnt unmarshal: %v", err.Error())
	}
	config.mapn = md.Next
	config.mapb = md.Previous
	results := md.Results
	for _, result := range results {
		fmt.Println(result.Name)
	}
	return nil
}
