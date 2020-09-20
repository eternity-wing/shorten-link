package main

import (
	"github.com/eternity-wing/short_link/database"
	"github.com/eternity-wing/short_link/handler/link"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:shorten", link.GetLink)
	app.Post("/api/v1/links", link.NewLink)
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
