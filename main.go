package main

import (
	"github.com/eternity-wing/short_link/controller/linkcontroller"
	"github.com/eternity-wing/short_link/database"
	"github.com/eternity-wing/short_link/model"
	"github.com/eternity-wing/short_link/repository/counterrepository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:shorten", linkcontroller.Get)
	app.Post("/api/v1/links", linkcontroller.New)
}

func main() {
	app := fiber.New()
	app.Use(recover.New())

	setupRoutes(app)
	loadEnvFile()
	database.InitiateMongo()
	loadFixtures()

	app.Listen(os.Getenv("PORT"))
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadFixtures() {
	ctrRepo := counterrepository.NewRepository()
	ctr := ctrRepo.Find(bson.M{"id": "link"})
	if ctr == nil {
		ctrRepo.Create(&model.Counter{ID: "link", Value: 0})
	}
}
