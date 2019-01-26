package core

import (
	"fmt"
	"github.com/Xarvalus/go-networking/webserver/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func Connect() *gorm.DB {
	pass := os.Getenv("DATABASE_PASSWORD")
	database := os.Getenv("DATABASE_NAME")

	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=127.0.0.1 port=5432 user=postgres dbname=%s password=%s sslmode=disable",
			database, pass))

	if err != nil {
		panic(err)
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Parent{}, &models.Children{})
}

func Close(db *gorm.DB) {
	err := db.Close()

	if err != nil {
		panic(err)
	}
}
