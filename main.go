package main

import (
	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/config"
	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/router"
)

var (
	logger *config.Logger
)

func main() {
	// Inicializa as variáveis de ambiente
	config.InitEnvVars()

	// Inicializa o logger
	logger = config.GetLogger("main")

	// Inicializa conexão com o banco de dados
	err := config.InitDB()
	if err != nil {
		logger.Errorf("Erro ao inicializar o banco de dados: %v", err)
		return
	}

	// Inicializa o router
	router.SetupRouter()
}
