package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	Port   string
	db     *gorm.DB
	logger *Logger
)

func InitEnvVars() error {
	// Initialize environment variables
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("Erro ao carregar variáveis de ambiente: %v", err)
	}

	Port = os.Getenv("API_PORT")
	return nil
}

func InitDB() error {
	var err error

	// Initialize database
	db, err = InitSQLite()
	if err != nil {
		return fmt.Errorf("Erro ao inicializar o banco de dados: %v", err)
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	// Initialize logger
	logger = NewLogger(prefix)
	return logger
}
