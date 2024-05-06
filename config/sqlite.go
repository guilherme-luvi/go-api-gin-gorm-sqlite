package config

import (
	"os"

	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// check if database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Arquivo do Banco de dados sqlite não encontrado, criando um novo...")

		// create db directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		// create db file
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
		logger.Info("Aquivo do banco de dados sqlite criado com sucesso")
	}

	// connect to sqlite database
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Erro ao conectar com o banco de dados sqlite: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Erro ao criar a tabela Opening: %v", err)
		return nil, err
	}

	logger.Info("Conexão com o banco de dados sqlite estabelecida com sucesso")
	return db, nil
}
