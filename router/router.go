package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	// Inicializa o router utilizando configurações padrão do Gin
	router := gin.Default()

	// Inicializa as rotas
	initalizeRoutes(router)

	// Inicializa o servidor na porta 5000
	router.Run(":5000")
}
