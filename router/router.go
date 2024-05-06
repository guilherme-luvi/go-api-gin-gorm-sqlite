package router

import "github.com/gin-gonic/gin"

func SetupRouter() {
	// Inicializa o router utilizando configurações padrão do Gin
	router := gin.Default()

	// Define uma rota para o método GET
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello!",
		})
	})

	router.Run(":5000")
}
