package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Println(pokemon.BaseExperience, randNum, threshold)
	if randNum > threshold {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	cfg.caughtPokemon[pokemonName] = pokemon

	return nil
}
