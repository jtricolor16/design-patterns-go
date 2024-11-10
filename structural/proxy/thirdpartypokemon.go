package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type libPokemon struct {
	httpInstance http.Client
}

func NewLibPokemon(httpInstance http.Client) PokemonIntegration {
	return &libPokemon{}
}

func (l *libPokemon) Show(name string) (Pokemon, error) {
	var pokemon Pokemon
	path := fmt.Sprintf("/api/v2/pokemon/%s", name)
	urlInstance := &url.URL{Scheme: "https", Host: "pokeapi.co", Path: path}
	fmt.Println(urlInstance.String())
	req := &http.Request{URL: urlInstance}
	response, err := l.httpInstance.Do(req)
	if err != nil {
		return pokemon, err
	}

	defer response.Body.Close()

	if response.StatusCode >= http.StatusOK && response.StatusCode <= http.StatusBadRequest {
		err = json.NewDecoder(response.Body).Decode(&pokemon)
		if err != nil {
			return pokemon, err
		}
		return pokemon, nil
	}

	return pokemon, fmt.Errorf("invalid request: %s", response.StatusCode)

}

func (l *libPokemon) ShowPokemon(names []string) ([]Pokemon, error) {
	return []Pokemon{}, nil
}
