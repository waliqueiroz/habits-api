package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/waliqueiroz/habits-api/internal/infra/repository/sqlite"
)

func main() {
	db, _ := sqlite.Connect()
	fmt.Println(db)
	app := fiber.New()

	app.Use(cors.New())

	app.Listen(":8080")
}
