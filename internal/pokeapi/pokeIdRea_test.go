package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"testing"
)

func TestPokeId(t *testing.T) {
	id := 1
	pokeEndpoint := "https://pokeapi.co/api/v2/location-area/"
	fullPokeUrl := pokeEndpoint + strconv.Itoa(id)

	fmt.Println(fullPokeUrl)
	req, err := http.NewRequest("GET", fullPokeUrl, nil)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
	c := Client{}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
	fmt.Println(data)
	pokemonEncounters := PokemonEncounters{}
	err = json.Unmarshal(data, &pokemonEncounters)

	fmt.Println(pokemonEncounters)

}
