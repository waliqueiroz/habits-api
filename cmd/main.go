package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/waliqueiroz/habits-api/internal/infra/repository/sqlite"
)

func main() {
	db, err := sqlite.Connect()
	if err != nil {
		panic(err)
	}

	err = sqlite.Migrate(db)
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Listen(":8080")
}
