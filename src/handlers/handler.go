package handlers

import (
	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/src/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("handler")
	db = config.GetDB()
}
