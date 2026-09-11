package main

import (
	"time"

	"github.com/JoshuaCarlton/pokedexcli/internal"
)

type config struct {
	commands map[string]cliCommand
	mapn     string
	mapb     string
	cache    internal.Cache
}

func getConfig() config {
	config := config{
		commands: getCommands(),
		mapn:     "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		mapb:     "",
		cache:    internal.NewCache(5 * time.Second),
	}
	return config
}
