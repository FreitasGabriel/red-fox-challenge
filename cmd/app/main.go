package main

import (
	"log"

	"github.com/FreitasGabriel/red-fox-challenge/config"
	"github.com/FreitasGabriel/red-fox-challenge/config/database/connection"
	"github.com/FreitasGabriel/red-fox-challenge/config/dependencies"
	"github.com/FreitasGabriel/red-fox-challenge/config/logger"
	"github.com/FreitasGabriel/red-fox-challenge/internal/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	c := gin.Default()

	config, err := config.LoadConfig(".")
	if err != nil {
		logger.Error("error to initialize the env configs", err)
		return
	}

	db, err := connection.NewConnectionDatabase(config)
	if err != nil {
		logger.Error("error to get connection database", err)
		return
	}

	handler := dependencies.InitDependencies(db)

	routes.InitRoutes(&c.RouterGroup, handler)

	err = c.Run(":8080")
	if err != nil {
		log.Fatal("error to start server")
	}
}
