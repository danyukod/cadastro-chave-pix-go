package main

import (
	"database/sql"
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/routes"
	"github.com/danyukod/cadastro-chave-pix-go/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	migrate_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

func main() {
	logger.Info("About to start PixKey API...")

	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

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

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	pixKeyController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, pixKeyController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
