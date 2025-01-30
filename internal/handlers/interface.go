package handlers

import (
	"github.com/FreitasGabriel/red-fox-challenge/internal/service"
	"github.com/gin-gonic/gin"
)

func NewPokemonHandler(service service.PokemonService) PokemonHandler {
	return &pokemonHandler{
		service,
	}
}

type pokemonHandler struct {
	service service.PokemonService
}

type PokemonHandler interface {
	GetPokemonByID(c *gin.Context)
}
