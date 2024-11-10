package main

import (
	"fmt"
	"net/http"
)

func main() {
	clientHttp := http.Client{}
	service := NewLibPokemon(clientHttp)
	serviceProxy := NewPokemonProxy(service)
	resp, err := serviceProxy.ShowPokemon([]string{"ditto", "pikachu", "squirtle", "jeff", "charmander"})
	if err == nil {
		fmt.Println(resp)
	}
	fmt.Println(err)
}
