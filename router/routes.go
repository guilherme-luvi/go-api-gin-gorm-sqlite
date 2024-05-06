package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/handlers"
)

func initalizeRoutes(router *gin.Engine) {
	// Initialize the handlers
	handlers.InitHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handlers.GetOpeningById)
		v1.POST("/opening", handlers.CreateOpening)
		v1.DELETE("/opening", handlers.DeleteOpening)
		v1.PUT("/opening", handlers.UpdateOpening)
		v1.GET("/openings", handlers.ListOpenings)
	}
}
