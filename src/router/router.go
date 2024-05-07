package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/src/config"
)

func SetupRouter() {
	// Inicializa o router utilizando configurações padrão do Gin
	router := gin.Default()

	// Inicializa as rotas
	initalizeRoutes(router)

	// Inicializa o servidor na porta 5000
	router.Run(":" + config.Port)
}
