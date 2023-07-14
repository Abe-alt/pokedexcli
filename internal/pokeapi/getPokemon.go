package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemons, error) {
	pokemonEndpoint := "https://pokeapi.co/api/v2/pokemon/"
	fullUrl := pokemonEndpoint + pokemonName

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Pokemons{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemons{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return Pokemons{}, nil
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemons{}, err
	}

	pokemons := Pokemons{}
	err = json.Unmarshal(data, &pokemons)
	if err != nil {
		return Pokemons{}, err
	}
	return pokemons, nil

}
