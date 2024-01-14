package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Check cache
    dat, ok := c.cache.Get(url)
	if ok {
		// Cache exists
		fmt.Println(".................Cache..................")
		locationsResp := LocationAreaResponse{}
	    err := json.Unmarshal(dat, &locationsResp)
	    if err != nil {
		    return LocationAreaResponse{}, err
	}
	    return locationsResp, nil
	}
	fmt.Println(".................No Cache..................")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationsResp := LocationAreaResponse{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	// Save to cache
	c.cache.Add(url, dat)
	return locationsResp, nil
}