package service

import "github.com/FreitasGabriel/red-fox-challenge/internal/repository"

func NewPokemonService(repository repository.PokemonRepository) PokemonService {
	return &pokemonService{
		repository,
	}
}

type pokemonService struct {
	repository repository.PokemonRepository
}

type PokemonService interface {
	GetPokemonByID() string
}
