package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waliqueiroz/habits-api/internal/application"
)

type HabitController struct {
	habitService application.HabitService
}

func NewHabitController(habitService application.HabitService) *HabitController {
	return &HabitController{
		habitService: habitService,
	}
}

func (c *HabitController) Create(ctx *fiber.Ctx) error {
	var habit HabitRequest

	if err := ctx.BodyParser(&habit); err != nil {
		return err
	}

	newHabit, err := c.habitService.Create(ctx.Context(), mapHabitDTOToDomain(habit))
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(mapHabitResponseFromDomain(*newHabit))
}
