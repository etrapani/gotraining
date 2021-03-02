package pokeapi

import (
	"encoding/json"
	"example.com/codelytv/cobra-03/cmd/ports"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	pokeAPIURL          = "https://pokeapi.co/api/v2"
)

type pokeApiRepo struct {
	url string
}

// NewPokeApiRepository fetch pokemons from https://pokeapi.co/api/v2
func NewPokeApiRepository() ports.GetPokemonsRepo {
	return &pokeApiRepo{url: pokeAPIURL}
}

func (p *pokeApiRepo) Execute(limit int, offset int) (pokemons []ports.Pokemon, err error) {
	response, err := http.Get(fmt.Sprintf("%v/pokemon?limit=%v&offset=%v", p.url, limit, offset))
	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var pokemonPage pagePokemonJson
	err = json.Unmarshal(contents, &pokemonPage)
	if err != nil {
		return nil, err
	}
	pokemons, err = p.getPokemonsInfo(pokemonPage.Results)
	if err != nil {
		return nil, err
	}
	return
}

func (p *pokeApiRepo) getPokemonsInfo(pagePokemonResults []pagePokemonResultJson) (pokemons []ports.Pokemon, err error) {
	for _, pokemon := range pagePokemonResults {
		response, err := http.Get(pokemon.Url)
		if err != nil {
			return nil, err
		}

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		var pokemonInfo pokemonInfoJson
		err = json.Unmarshal(contents, &pokemonInfo)
		if err != nil {
			return nil, err
		}
		pokemons = append(pokemons, toPokemon(pokemonInfo))
	}
	return
}

func toPokemon(source pokemonInfoJson) (pokemon ports.Pokemon) {
	pokemon.Number = source.Id
	pokemon.Name = source.Name
	pokemon.Height = source.Height
	pokemon.Weight = source.Weight
	return
}

type pagePokemonJson struct {
	Count    int                     `json:"count"`
	Next     string                  `json:"next"`
	Previous string                  `json:"previous"`
	Results  []pagePokemonResultJson `json:"results"`
}

type pagePokemonResultJson struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type pokemonInfoJson struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
}
