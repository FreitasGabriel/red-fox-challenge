package routes

import "github.com/gin-gonic/gin"

func InitRoutes(cr *gin.RouterGroup) {
	cr.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, "checked")
	})
}
