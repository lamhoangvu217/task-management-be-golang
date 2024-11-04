package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/lamhoangvu217/task-management-be-golang/database"
	"github.com/lamhoangvu217/task-management-be-golang/routes"
	"log"
	"os"
)

func main() {
	database.Connect()
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
	port := os.Getenv("PORT")
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, http://localhost:3000", // Replace with the actual frontend origin(s)
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowCredentials: true,
	}))
	routes.Setup(app)
	app.Listen(":" + port)

}
