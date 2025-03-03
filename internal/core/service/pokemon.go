package service

import (
	"github.com/FreitasGabriel/red-fox-challenge/config/logger"
	"github.com/FreitasGabriel/red-fox-challenge/internal/core/domain"
)

func (ps *pokemonService) GetPokemonByID(id string) (*domain.Pokemon, error) {

	pokemon, err := ps.repository.GetPokemonByID(id)
	if err != nil {
		logger.Error("error to get pokemon by id", err)
		return nil, err
	}

	return pokemon, nil
}

func (ps *pokemonService) CreatePokemonInBatch(pokemon []*domain.Pokemon, batchSize int) error {

	err := ps.repository.InsertPokemonInBatch(pokemon, batchSize)
	if err != nil {
		logger.Error("error to create pokemon in batch", err)
		return err
	}

	return nil
}
