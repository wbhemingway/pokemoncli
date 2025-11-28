package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationAreaResponse struct {
	Count   int     `json:"count"`
	Next    *string `json:"next"`
	Prev    *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(urlPointer *string) ([]string, *string, *string, error) {
	var url string
	if urlPointer == nil {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = *urlPointer
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error getting response from url: %w", err)
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	var data locationAreaResponse
	err = decoder.Decode(&data)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error decoding response body: %w", err)
	}

	names := make([]string, 0, len(data.Results))
	for _, res := range data.Results {
		names = append(names, res.Name)
	}
	return names, data.Prev, data.Next, nil
}
