package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
