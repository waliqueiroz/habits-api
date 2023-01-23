package rest

import (
	"time"

	"github.com/waliqueiroz/habits-api/internal/domain"
)

type HabitRequest struct {
	Title    string `json:"title"`
	Weekdays []int  `json:"weekdays"`
}

func mapHabitDTOToDomain(habit HabitRequest) domain.Habit {
	return domain.NewHabit(habit.Title, habit.Weekdays)
}

type HabitResponse struct {
	ID        string                 `json:"id"`
	Title     string                 `json:"title"`
	Weekdays  []HabitWeekdayResponse `json:"weekdays"`
	CreatedAt time.Time              `json:"created_at"`
}

type HabitWeekdayResponse struct {
	ID      string `json:"id"`
	HabitID string `json:"habit_id"`
	Weekday int    `json:"weekday"`
}

func mapHabitResponseFromDomain(habit domain.Habit) HabitResponse {
	habitWeekdays := make([]HabitWeekdayResponse, len(habit.Weekdays))

	for key, value := range habit.Weekdays {
		habitWeekdays[key] = HabitWeekdayResponse{
			ID:      value.ID,
			HabitID: value.HabitID,
			Weekday: value.Weekday,
		}
	}

	return HabitResponse{
		ID:        habit.ID,
		Title:     habit.Title,
		Weekdays:  habitWeekdays,
		CreatedAt: habit.CreatedAt,
	}
}
