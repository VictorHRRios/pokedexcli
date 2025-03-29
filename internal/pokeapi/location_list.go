package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/VictorHRRios/pokedexcli/internal/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2"

func ListLocations(pageUrl *string, cache *pokecache.Cache) (LocationArea, error) {
	fullUrl := baseUrl + "/location-area"
	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	if cached, ok := cache.Get(fullUrl); ok {
		fmt.Println("Using cacheed")
		locationArea := LocationArea{}
		err := json.Unmarshal(cached, &locationArea)
		if err != nil {
			return LocationArea{}, fmt.Errorf("Cannot unmarshal")
		}
		return locationArea, nil
	}

	res, err := http.Get(fullUrl)
	if err != nil {
		return LocationArea{}, fmt.Errorf("Cannot fetch api")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return LocationArea{}, fmt.Errorf("Response failed with status code %v", res.Status)
	}

	cache.Add(fullUrl, body)

	locationArea := LocationArea{}
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationArea{}, fmt.Errorf("Cannot unmarshal")
	}

	return locationArea, nil
}
