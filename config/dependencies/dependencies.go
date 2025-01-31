package dependencies

import (
	"github.com/FreitasGabriel/red-fox-challenge/internal/api/handlers"
	"github.com/FreitasGabriel/red-fox-challenge/internal/core/service"
	"github.com/FreitasGabriel/red-fox-challenge/internal/repository"
	"gorm.io/gorm"
)

func InitDependencies(db *gorm.DB) handlers.PokemonHandler {
	repo := repository.NewPokemonRepository(db)
	serv := service.NewPokemonService(repo)
	hand := handlers.NewPokemonHandler(serv)

	return hand
}
