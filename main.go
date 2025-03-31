package main

import "github.com/VictorHRRios/pokedexcli/internal/pokeapi"

func main() {
	cfg := &config{retr: pokeapi.GetRetrieve(5),
		pokedex: map[string]pokeapi.PokemonDetail{}}
	repl(cfg)
}
