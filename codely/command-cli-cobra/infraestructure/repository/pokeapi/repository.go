package pokeapi

import (
	"encoding/json"
	"example.com/codelytv/cobra-03/cmd/ports"
	"fmt"
	"io/ioutil"
	"net/http"
	jsoniter "github.com/json-iterator/go"
	"sync"
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
	err = p.betterUnmarshal(contents, &pokemonPage)
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
	mu := &sync.Mutex{}
	totalPokemons := len(pagePokemonResults)
	type pokePos struct {
		pokemon ports.Pokemon
		pos int
	}

	queue := make(chan ports.Pokemon)
	for i, pokemon := range pagePokemonResults {
		go p.fetchPokemonInfo(pokemon.Url, queue, i)
	}
	i:=0
	for i < totalPokemons {
		select {
			case pk := <- queue:
				mu.Lock()
				pokemons = append(pokemons, pk)
				mu.Unlock()
				i++
		}
	}
	return pokemons, nil
}

func (p *pokeApiRepo) fetchPokemonInfo(url string, queue chan ports.Pokemon, index int) {
	response, err := http.Get(url)
	if err != nil {
		return
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	var pokemonInfo pokemonInfoJson
	err = p.betterUnmarshal(contents, &pokemonInfo)
	if err != nil {
		return
	}

	queue <- toPokemon(pokemonInfo)
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

func (p *pokeApiRepo) standardUnmarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	return nil
}

func (p *pokeApiRepo) betterUnmarshal(data []byte, v interface{}) error {
	var js = jsoniter.ConfigCompatibleWithStandardLibrary

	err := js.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	return nil
}

