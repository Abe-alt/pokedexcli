package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasResp, error) {

	endpoint := "/location-area/"
	fullURL := baseUrl + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err // return 0 value Location resp
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}

	defer resp.Body.Close() // close resp object when we're done w/ it

	if resp.StatusCode > 399 { // the client will take care of redirect errors 300
		return LocationAreasResp{}, fmt.Errorf("bad status code %v ", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreaResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreaResp, nil

}
