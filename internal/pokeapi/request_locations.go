package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) RequestLocations(pageURL *string) (LocationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationsResp := LocationArea{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationArea{}, err
	}

	return locationsResp, nil
}
