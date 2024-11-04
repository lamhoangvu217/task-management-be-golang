package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lamhoangvu217/task-management-be-golang/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	log.Println("env", os.Getenv("ENV"))
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Fatal("Error loading .env.local file")
		}
	} else {
		err := godotenv.Load(".env.production")
		if err != nil {
			log.Fatal("Error loading .env.production file")
		}
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	} else {
		log.Println("Connected to the database")
	}
	DB = database
	database.AutoMigrate(
		&models.User{},
		&models.Task{},
		&models.Subtask{},
		&models.Label{},
	)
}
