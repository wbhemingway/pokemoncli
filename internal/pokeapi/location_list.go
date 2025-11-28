package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	var err error
	dat, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespLocations{}, fmt.Errorf("error creating request from url: %w", err)
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespLocations{}, fmt.Errorf("error getting response from url: %w", err)
		}

		defer resp.Body.Close()
		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespLocations{}, fmt.Errorf("error reading response from body: %w", err)
		}
		c.cache.Add(url, dat)
	}

	locationResp := RespLocations{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespLocations{}, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	return locationResp, nil
}
