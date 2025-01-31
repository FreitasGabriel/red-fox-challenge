package repository

import (
	"gorm.io/gorm"
)

func NewPokemonRepository(repo *gorm.DB) PokemonRepository {
	return &pokemonRepository{
		repository: repo,
	}
}

type pokemonRepository struct {
	repository *gorm.DB
}

type PokemonRepository interface {
	GetPokemonByID(id string) string
}
