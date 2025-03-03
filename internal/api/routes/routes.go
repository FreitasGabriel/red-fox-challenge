package routes

import (
	"github.com/FreitasGabriel/red-fox-challenge/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(cr *gin.RouterGroup, handler handlers.PokemonHandler) {
	cr.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, "checked")
	})

	cr.GET("/pokemon/:id", handler.GetPokemonByID)
	cr.POST(("/pokemon"), handler.CreatePokemonInBatch)
}
