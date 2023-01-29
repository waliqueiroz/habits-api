package rest

import (
	"time"

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
	var habit HabitRequestDTO

	if err := ctx.BodyParser(&habit); err != nil {
		return err
	}

	err := c.habitService.Create(ctx.Context(), mapHabitRequestToDomain(habit))
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

func (c *HabitController) GetDayResume(ctx *fiber.Ctx) error {
	dateString := ctx.Query("date")

	date, err := time.Parse(time.RFC3339Nano, dateString)
	if err != nil {
		return err
	}

	dayResume, err := c.habitService.GetDayResume(ctx.Context(), date)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(mapDayResumeFromDomain(*dayResume))
}

func (c *HabitController) ToggleHabit(ctx *fiber.Ctx) error {
	habitID := ctx.Params("habitID")

	err := c.habitService.ToggleHabit(ctx.Context(), habitID)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}
