package factories

import (
	"github.com/danyukod/cadastro-chave-pix-go/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDatabaseFactory(conf configs.Config) (*gorm.DB, error) {
	dsn := conf.GetUser() + ":" + conf.GetPassword() + "@tcp(" + conf.GetHost() + ":" + conf.GetPort() + ")/" + conf.GetName() + "?charset=utf8mb4&parseTime=True&loc=Local"

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
