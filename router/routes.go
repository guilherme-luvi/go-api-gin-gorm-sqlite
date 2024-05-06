package router

import "github.com/gin-gonic/gin"

func initalizeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "GET /api/v1/opening",
			})
		})
		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "POST /api/v1/opening",
			})
		})
		v1.DELETE("/opening", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "DELETE /api/v1/opening",
			})
		})
		v1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "PUT /api/v1/opening",
			})
		})
		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "GET ALL /api/v1/openings",
			})
		})
	}
}
