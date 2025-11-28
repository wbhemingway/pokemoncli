package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocations{}, fmt.Errorf("error creating request from url: %w", err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, fmt.Errorf("error getting response from url: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return RespLocations{}, fmt.Errorf("unexpected status from get request: %v", resp.Status)
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	var data RespLocations
	err = decoder.Decode(&data)
	if err != nil {
		return RespLocations{}, fmt.Errorf("error decoding response body: %w", err)
	}

	return data, nil
}
