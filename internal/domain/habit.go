package domain

import (
	"context"
	"time"

	"github.com/waliqueiroz/habits-api/pkg/identity"
)

type HabitRepository interface {
	Create(ctx context.Context, habit Habit) (*Habit, error)
}

type Habit struct {
	ID        string
	Title     string
	Weekdays  []HabitWeekday
	CreatedAt time.Time
}

type HabitWeekday struct {
	ID      string
	HabitID string
	Weekday int
}

func NewHabit(title string, weekdays []int) Habit {
	habitID := identity.NewULID()

	habitWeekdays := make([]HabitWeekday, len(weekdays))

	for i, value := range weekdays {
		habitWeekdays[i] = HabitWeekday{
			ID:      identity.NewULID(),
			HabitID: habitID,
			Weekday: value,
		}
	}

	return Habit{
		ID:        habitID,
		Title:     title,
		Weekdays:  habitWeekdays,
		CreatedAt: truncateToDay(time.Now()).UTC(),
	}
}

func truncateToDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
