package factories

import (
	"database/sql"
	"github.com/danyukod/cadastro-chave-pix-go/configs"
	"github.com/golang-migrate/migrate/v4"
	migrate_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewPixKeyDatabaseFactory(dbConf configs.DbConfig) (*gorm.DB, error) {
	dsn := dbConf.User + ":" + dbConf.Password + "@tcp(" + dbConf.Host + ":" + dbConf.Port + ")/" + dbConf.Name + "?charset=utf8mb4&parseTime=True&loc=Local"

	if dbConf.MigrationTag {
		db, _ := sql.Open(dbConf.Driver, dsn)
		driver, _ := migrate_mysql.WithInstance(db, &migrate_mysql.Config{})
		m, _ := migrate.NewWithDatabaseInstance(dbConf.MigrationSource, dbConf.Name, driver)
		m.Up()
	}

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
