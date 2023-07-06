package factories

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
	migrationSource := os.Getenv("MIGRATION_SOURCE")
	dbName := os.Getenv("DATABASE_NAME")
	dbDriver := os.Getenv("DATABASE_DRIVER")

	if migrationTag == "ON" {
		db, _ := sql.Open(dbDriver, dsn)
		driver, _ := migrate_mysql.WithInstance(db, &migrate_mysql.Config{})
		m, _ := migrate.NewWithDatabaseInstance(migrationSource, dbName, driver)
		m.Up()
	}

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
