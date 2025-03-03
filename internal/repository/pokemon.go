package repository

import (
	"github.com/FreitasGabriel/red-fox-challenge/config/logger"
	"github.com/FreitasGabriel/red-fox-challenge/internal/core/domain"
	"gorm.io/gorm"
)

func (pr *pokemonRepository) GetPokemonByID(id string) (*domain.Pokemon, error) {

	pokemon := &domain.Pokemon{}

	pr.repository.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(pokemon, id).Error; err != nil {
			logger.Error("error to get pokemon by id", err)
			return err
		}
		return nil
	})

	return pokemon, nil
}

func (pr *pokemonRepository) InsertPokemonInBatch(pokemon []*domain.Pokemon, batchSize int) error {

	pr.repository.Transaction(func(tx *gorm.DB) error {
		logger.Info("insert pokemons in batch into databse...")
		if err := tx.CreateInBatches(pokemon, batchSize).Error; err != nil {
			logger.Error("error to insert pokemons into database, doing rollback...", err)
			tx.Rollback()
			return err
		}
		return nil
	})

	return nil
}
