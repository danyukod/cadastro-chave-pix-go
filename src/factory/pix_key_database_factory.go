package factory

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	migrate_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func NewPixKeyDatabaseFactory() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_CONNECTION_STRING")
	migrationTag := os.Getenv("MIGRATION_TAG")

	if migrationTag == "ON" {
		db, _ := sql.Open("mysql", dsn)
		driver, _ := migrate_mysql.WithInstance(db, &migrate_mysql.Config{})
		m, _ := migrate.NewWithDatabaseInstance(
			"file:../migrations",
			"cadastro_chave_pix",
			driver,
		)

		m.Up()
	}

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
