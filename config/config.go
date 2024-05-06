package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

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
