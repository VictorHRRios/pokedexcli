package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/VictorHRRios/pokedexcli/internal/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Retrieve struct {
	cache *pokecache.Cache
}

func GetRetrieve(timer int) *Retrieve {
	c := pokecache.NewCache(time.Duration(timer * int(time.Minute)))
	return &Retrieve{
		cache: &c,
	}
}

func (retr Retrieve) ListLocations(pageUrl *string) (LocationArea, error) {
	fullUrl := baseUrl + "/location-area" + "?offset=0&limit=20"
	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	if cached, ok := retr.cache.Get(fullUrl); ok {
		//fmt.Println("Using cache'd")
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

	retr.cache.Add(fullUrl, body)

	locationArea := LocationArea{}
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationArea{}, fmt.Errorf("Cannot unmarshal")
	}

	return locationArea, nil
}
