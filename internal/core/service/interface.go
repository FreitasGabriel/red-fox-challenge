package service

import (
	"github.com/FreitasGabriel/red-fox-challenge/internal/core/domain"
	"github.com/FreitasGabriel/red-fox-challenge/internal/repository"
)

func NewPokemonService(repository repository.PokemonRepository) PokemonService {
	return &pokemonService{
		repository,
	}
}

type pokemonService struct {
	repository repository.PokemonRepository
}

type PokemonService interface {
	GetPokemonByID(id string) (*domain.Pokemon, error)
	ReadFile(filepath string) ([]*domain.Pokemon, error)
	CreatePokemonInBatch(pokemons []*domain.Pokemon, batchSize int) error
}
