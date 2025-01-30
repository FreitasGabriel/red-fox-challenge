package main

import (
	"log"

	"github.com/FreitasGabriel/red-fox-challenge/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	c := gin.Default()

	routes.InitRoutes(&c.RouterGroup)

	err := c.Run(":8080")
	if err != nil {
		log.Fatal("error to start server")
	}
}
