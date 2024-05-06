package config

import "gorm.io/gorm"

var (
	db     *gorm.DB
	logger *Logger
)

func InitDB() error {
	return nil
}

func InitLogger(prefix string) *Logger {
	// Initialize logger
	logger = NewLogger(prefix)
	return logger
}
