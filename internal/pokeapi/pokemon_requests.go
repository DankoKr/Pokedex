package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	// Check cache
	dat, ok := c.cache.Get(url)
	if ok {
		// Cache exists
		fmt.Println(".................Cache..................")
		pokemonResp := Pokemon{}
		err := json.Unmarshal(dat, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return Pokemon{}, nil
	}
	fmt.Println(".................No Cache..................")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	// Save to cache
	c.cache.Add(url, dat)
	return pokemonResp, nil
}