package ports

// Beer representation of beer into data struct
type Pokemon struct {
	Number int
	Name   string
	Height int
	Weight int
}

// Get pokemons
type GetPokemonsRepo interface {
	Execute(limit int, offset int) ([]Pokemon, error)
}

// Get pokemons
type SavePokemonsInFileRepo interface {
	Execute(pokemons []Pokemon, path string, filename string) error
}
