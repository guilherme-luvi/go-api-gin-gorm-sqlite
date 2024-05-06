package handlers

import "github.com/gin-gonic/gin"

func CreateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "POST /api/v1/opening",
	})
}
