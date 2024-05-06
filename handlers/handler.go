package handlers

import "github.com/gin-gonic/gin"

func CreateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "POST /api/v1/opening",
	})
}

func ShowOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "GET /api/v1/opening",
	})
}

func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "DELETE /api/v1/opening",
	})
}

func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "PUT /api/v1/opening",
	})
}

func ListOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "GET /api/v1/openings",
	})
}
