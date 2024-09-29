package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ahmed-afzal1/restaurant/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database! %v", err)
	}

	err = database.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.SubCategory{},
		&models.Cuisine{},
		&models.Restaurant{},
		&models.Addon{},
		&models.Food{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate to database! %v", err)

	}

	DB = database
}
