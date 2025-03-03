package handlers

import (
	"path/filepath"

	"github.com/FreitasGabriel/red-fox-challenge/config/logger"
	"github.com/FreitasGabriel/red-fox-challenge/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (ph *pokemonHandler) GetPokemonByID(c *gin.Context) {

	file, err := filepath.Abs("../../assets/pokemon-go.xlsx")
	if err != nil {
		logger.Error("error to get path", err)
		return
	}

	ph.service.ReadFile(file)

	c.JSON(200, file)
}

func (ph *pokemonHandler) CreatePokemonInBatch(c *gin.Context) {

	var pokemon []*domain.Pokemon

	file, err := filepath.Abs("../../assets/pokemon-go.xlsx")
	if err != nil {
		logger.Error("error to get path", err)
		return
	}

	pokemon, err = ph.service.ReadFile(file)
	if err != nil {
		logger.Error("error to read file", err)
		return
	}

	err = ph.service.CreatePokemonInBatch(pokemon, 50)
	if err != nil {
		logger.Error("error to create pokemon in batch", err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, "pokemons inserted in database successfully")
}
