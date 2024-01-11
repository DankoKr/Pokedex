package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations() (LocationAreaResponse, error) {
	url := baseURL + "/location-area"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationsResp := LocationAreaResponse{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationsResp, nil
}