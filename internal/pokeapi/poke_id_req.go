package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (c *Client) ListPokeId() (PokemonEncounters, error) {

	var id int
	pokeEndpoint := "https://pokeapi.co/api/v2/location-area/"
	fullPokeUrl := pokeEndpoint + strconv.Itoa(id)

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
