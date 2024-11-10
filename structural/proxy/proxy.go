package main

import (
	"fmt"
	"sync"
)

type pokemonProxy struct {
	integration PokemonIntegration
}

func NewPokemonProxy(integration PokemonIntegration) PokemonIntegration {
	return &pokemonProxy{
		integration: integration,
	}
}

func (p *pokemonProxy) Show(name string) (Pokemon, error) {
	return p.integration.Show(name)
}

func (p *pokemonProxy) ShowPokemon(names []string) ([]Pokemon, error) {
	var pokemon []Pokemon
	var wg sync.WaitGroup

	for _, name := range names {
		wg.Add(1)

		go func(name string) {
			defer wg.Done()
			resp, err := p.integration.Show(name)
			if err != nil {
				fmt.Printf("error to get pokemon %s\n", name)
				return
			}
			fmt.Println(resp)
			pokemon = append(pokemon, resp)
		}(name)

	}

	wg.Wait()

	return pokemon, nil
}
