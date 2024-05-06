package handlers

import "github.com/gin-gonic/gin"

func sendError(ctx *gin.Context, statusCode int, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"error":      message,
	})
}

func sendSuccess(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"data":       data,
	})
}
