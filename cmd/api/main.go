package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/waliqueiroz/habits-api/internal/application"
	"github.com/waliqueiroz/habits-api/internal/infra/entrypoint"
	"github.com/waliqueiroz/habits-api/internal/infra/entrypoint/rest"
	"github.com/waliqueiroz/habits-api/internal/infra/repository/sqlite"
)

func main() {
	db, err := sqlite.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = sqlite.Migrate(db.DB)
	if err != nil {
		panic(err)
	}

	if shouldSeedDB() {
		sqlite.Seed(db.DB)
	}

	app := fiber.New()

	app.Use(cors.New())

	habitRepository := sqlite.NewHabitRepository(db)
	habitService := application.NewHabitService(habitRepository)
	habitController := rest.NewHabitController(habitService)

	entrypoint.CreateRoutes(app, habitController)

	app.Listen(":8080")
}

func shouldSeedDB() bool {
	flag.Parse()
	args := flag.Args()

	return len(args) >= 1 && args[0] == "seed"
}
