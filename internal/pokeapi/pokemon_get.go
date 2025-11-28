package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + name
	var err error
	dat, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespPokemon{}, fmt.Errorf("error creating request from url: %w", err)
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespPokemon{}, fmt.Errorf("error getting response from url: %w", err)
		}

		if resp.StatusCode != http.StatusOK {
			return RespPokemon{}, fmt.Errorf("unexpected status code received from url:%v, %v", url, resp.Status)
		}

		defer resp.Body.Close()
		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespPokemon{}, fmt.Errorf("error reading response from body: %w", err)
		}

		c.cache.Add(url, dat)
	}

	pokemonResp := RespPokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokemon{}, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	return pokemonResp, nil
}
