package csv

import (
	"encoding/csv"
	"example.com/gotraining/command-cli-cobra/cmd/ports"
	"fmt"
	"log"
	"os"
	"strconv"
)

type pokeApiRepo struct {
}

func NewPokeApiRepository() ports.SavePokemonsInFileRepo {
	return &pokeApiRepo{}
}

func (p *pokeApiRepo) Execute(pokemons []ports.Pokemon, path string, filename string) (err error) {
	f, err := os.Create(fmt.Sprintf("%v/%v.csv", path, filename))
	defer f.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()
	w.Write([]string{"Number", "name", "weight", "height"})
	for _, pokemon := range pokemons {
		if err := w.Write(parseTo(pokemon)); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	return
}

func parseTo(pokemon ports.Pokemon) []string {
	return []string{strconv.Itoa(pokemon.Number), pokemon.Name, strconv.Itoa(pokemon.Weight), strconv.Itoa(pokemon.Height)}
}

