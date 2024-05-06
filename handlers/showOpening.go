package handlers

import "github.com/gin-gonic/gin"

func ShowOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "GET /api/v1/opening",
	})
}
