package factories

import (
	"github.com/danyukod/cadastro-chave-pix-go/configs"
	"github.com/danyukod/cadastro-chave-pix-go/configs/logger"
	"gorm.io/gorm"
)

func NewPixKeyDatabaseFactory(conf *configs.Config) (*gorm.DB, error) {
	if err := configs.MigrateDatabase(*conf); err != nil {
		logger.Error("Database migration failed: ", err)
		return nil, err
	}

	database, err := NewGormDatabaseFactory(*conf)
	if err != nil {
		logger.Error("Failed to create Gorm Database: ", err)
		return nil, err
	}

	return database, nil
}
