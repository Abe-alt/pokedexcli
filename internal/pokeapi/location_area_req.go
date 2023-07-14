package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResp, error) {

	endpoint := "/location-area/"
	fullURL := baseUrl + endpoint

	if pageUrl != nil {
		fullURL = *pageUrl
	}

	data, ok := c.cache.Get(fullURL)

	if ok {
		//fmt.Println("got from cache")
		locationAreaResp := LocationAreasResp{}
		err := json.Unmarshal(data, &locationAreaResp)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return locationAreaResp, nil

	}

	//fmt.Println("got from req")
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

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullURL, data) // add to the cache

	locationAreaResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreaResp, nil

}

func (c *Client) GetLocationArea(locationName string) (PokemonEncounters, error) {

	pokeEndpoint := "https://pokeapi.co/api/v2/location-area/"
	fullPokeUrl := pokeEndpoint + locationName

	req, err := http.NewRequest("GET", fullPokeUrl, nil)
	if err != nil {
		return PokemonEncounters{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonEncounters{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return PokemonEncounters{}, fmt.Errorf("bad status code,%s", err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonEncounters{}, nil
	}
	pokemonEncounters := PokemonEncounters{}
	err = json.Unmarshal(data, &pokemonEncounters)
	if err != nil {
		return PokemonEncounters{}, nil
	}

	return pokemonEncounters, nil

}
