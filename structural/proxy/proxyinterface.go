package main

type Pokemon struct {
	Name string `json:"name,omitempty"`
	Id   int    `json:"id,omitempty"`
}

type PokemonIntegration interface {
	Show(name string) (Pokemon, error)
	ShowPokemon(names []string) ([]Pokemon, error)
}
