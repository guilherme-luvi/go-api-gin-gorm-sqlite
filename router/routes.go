package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/handlers"
)

func initalizeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handlers.ShowOpeningHandler)
		v1.POST("/opening", handlers.CreateOpeningHandler)
		v1.DELETE("/opening", handlers.DeleteOpeningHandler)
		v1.PUT("/opening", handlers.UpdateOpeningHandler)
		v1.GET("/openings", handlers.ListOpeningsHandler)
	}
}
