package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (retr Retrieve) GetPokemon(pokemonID *string) (PokemonDetail, error) {
	fullUrl := baseUrl + "/pokemon"
	if pokemonID == nil {
		return PokemonDetail{}, fmt.Errorf("Nothing to display")
	}
	fullUrl += "/" + *pokemonID

	if cached, ok := retr.cache.Get(fullUrl); ok {
		pokemonDetail := PokemonDetail{}
		err := json.Unmarshal(cached, &pokemonDetail)
		if err != nil {
			return PokemonDetail{}, fmt.Errorf("Cannot unmarshal")
		}
		return pokemonDetail, nil
	}

	res, err := http.Get(fullUrl)
	if err != nil {
		return PokemonDetail{}, fmt.Errorf("Cannot fetch api")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return PokemonDetail{}, fmt.Errorf("Response failed with status code %v", res.Status)
	}

	retr.cache.Add(fullUrl, body)

	pokemonDetail := PokemonDetail{}
	err = json.Unmarshal(body, &pokemonDetail)
	if err != nil {
		return PokemonDetail{}, fmt.Errorf("Cannot unmarshal")
	}

	return pokemonDetail, nil
}
