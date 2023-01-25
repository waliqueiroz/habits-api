package entrypoint

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waliqueiroz/habits-api/internal/infra/entrypoint/rest"
)

func CreateRoutes(router fiber.Router, habitController *rest.HabitController) {
	api := router.Group("/api")

	api.Post("/habits", habitController.Create)
	api.Get("/day", habitController.GetDayResume)
}
