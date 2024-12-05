package main

import "github.com/katsuikeda/pokedexcli/internal/pokeapi"

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)
}
