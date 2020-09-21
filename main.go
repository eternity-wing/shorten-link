package main

import (
	"github.com/eternity-wing/short_link/controller/linkcontroller"
	"github.com/eternity-wing/short_link/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:shorten", linkcontroller.Get)
	app.Post("/api/v1/links", linkcontroller.New)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	loadEnvFile()
	database.InitiateMongo()

	app.Listen(os.Getenv("PORT"))
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
