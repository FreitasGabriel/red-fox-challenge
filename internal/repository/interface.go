package repository

import "database/sql"

func NewPokemonRepository(repo sql.DB) PokemonRepository {
	return &pokemonRepository{
		repository: repo,
	}
}

type pokemonRepository struct {
	repository sql.DB
}

type PokemonRepository interface {
	GetPokemonByID(id string) string
}
