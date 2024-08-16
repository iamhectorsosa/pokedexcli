package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]

	if ok {
		return fmt.Errorf("%s has already been caught", pokemon.Name)
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		return err
	}

	randNum := rand.Intn(pokemon.BaseExperience)

	const threshold = 50
	fmt.Println(pokemon.BaseExperience, randNum, threshold)
	if randNum > threshold {
		return fmt.Errorf("Failed to catch %s", pokemon.Name)
	}

	cfg.caughtPokemon[pokemon.Name] = pokemon
	fmt.Printf("%s was caught\n", pokemon.Name)
	return nil
}
