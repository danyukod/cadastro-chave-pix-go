package configs

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	migrate_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
)

func MigrateDatabase(conf Config) error {
	dsn := conf.GetUser() + ":" + conf.GetPassword() + "@tcp(" + conf.GetHost() + ":" + conf.GetPort() + ")/" + conf.GetName() + "?charset=utf8mb4&parseTime=True&loc=Local"

	if conf.GetMigrationTag() {
		db, _ := sql.Open(conf.GetDriver(), dsn)
		driver, _ := migrate_mysql.WithInstance(db, &migrate_mysql.Config{})
		m, _ := migrate.NewWithDatabaseInstance("file://migrations", conf.GetName(), driver)
		m.Up()
	}
	return nil
}
