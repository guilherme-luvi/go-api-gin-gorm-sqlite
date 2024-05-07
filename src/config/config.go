package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	Port      string
	SecretKey []byte
	db        *gorm.DB
	logger    *Logger
)

func InitEnvVars() error {
	// Initialize environment variables
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("erro ao carregar variáveis de ambiente: %v", err)
	}

	Port = os.Getenv("API_PORT")
	SecretKey = []byte(os.Getenv("SECRET_KEY"))

	fmt.Print("variáveis de ambiente carregadas com sucesso\n")
	return nil
}

func InitDB() error {
	var err error

	// Initialize database
	db, err = InitSQLite()
	if err != nil {
		return fmt.Errorf("erro ao inicializar o banco de dados: %v", err)
	}
	fmt.Print("banco de dados inicializado com sucesso\n")
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
