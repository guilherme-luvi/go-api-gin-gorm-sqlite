package main

import (
	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/config"
	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/router"
)

func main() {
	// Inicializa o logger
	logger := config.GetLogger("main")

	// Inicializa as variáveis de ambiente
	err := config.InitEnvVars()
	if err != nil {
		logger.Errorf("Erro ao inicializar as variáveis de ambiente: %v", err)
		return

	}

	// Inicializa conexão com o banco de dados
	err = config.InitDB()
	if err != nil {
		logger.Errorf("Erro ao inicializar o banco de dados: %v", err)
		return
	}

	// Inicializa o router
	router.SetupRouter()
}
