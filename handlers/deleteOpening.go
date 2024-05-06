package handlers

import "github.com/gin-gonic/gin"

func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "DELETE /api/v1/opening",
	})
}
