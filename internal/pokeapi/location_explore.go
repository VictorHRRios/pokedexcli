package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (retr Retrieve) ExploreLoaction(locationID *string) (LocationDetail, error) {
	fullUrl := baseUrl + "/location-area"
	if locationID == nil {
		return LocationDetail{}, fmt.Errorf("Nothing to display")
	}
	fullUrl += "/" + *locationID

	if cached, ok := retr.cache.Get(fullUrl); ok {
		//fmt.Println("Using cache'd")
		locationArea := LocationDetail{}
		err := json.Unmarshal(cached, &locationArea)
		if err != nil {
			return LocationDetail{}, fmt.Errorf("Cannot unmarshal")
		}
		return locationArea, nil
	}

	res, err := http.Get(fullUrl)
	if err != nil {
		return LocationDetail{}, fmt.Errorf("Cannot fetch api")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return LocationDetail{}, fmt.Errorf("Response failed with status code %v", res.Status)
	}

	retr.cache.Add(fullUrl, body)

	locationDetail := LocationDetail{}
	err = json.Unmarshal(body, &locationDetail)
	if err != nil {
		return LocationDetail{}, fmt.Errorf("Cannot unmarshal")
	}

	return locationDetail, nil
}
