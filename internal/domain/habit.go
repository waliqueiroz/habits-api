package domain

import (
	"context"
	"time"

	"github.com/waliqueiroz/habits-api/pkg/identity"
)

type HabitRepository interface {
	Create(ctx context.Context, habit Habit) error
	GetPossibleHabits(ctx context.Context, date time.Time) ([]Habit, error)
	GetCompletedHabits(ctx context.Context, date time.Time) ([]Habit, error)
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
		CreatedAt: truncateToDay(time.Now()),
	}
}

func truncateToDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

type DayResume struct {
	PossibleHabits  []Habit
	CompletedHabits []Habit
}
