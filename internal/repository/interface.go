package repository

import (
	"github.com/FreitasGabriel/red-fox-challenge/internal/core/domain"
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
	GetPokemonByID(id string) (*domain.Pokemon, error)
	InsertPokemonInBatch(pokemon []*domain.Pokemon, batchSize int) error
}
