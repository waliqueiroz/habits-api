package domain

import (
	"context"

	"github.com/waliqueiroz/habits-api/pkg/identity"
)

type DayHabitRepository interface {
	Create(ctx context.Context, dayHabit DayHabit) error
	ExistsByDayIDAndHabitID(ctx context.Context, dayID string, habitID string) (bool, error)
	DeleteByDayIDAndHabitID(ctx context.Context, dayID string, habitID string) error
}

type DayHabit struct {
	ID      string
	DayID   string
	HabitID string
}

func NewDayHabit(dayID string, habitID string) DayHabit {
	return DayHabit{
		ID:      identity.NewULID(),
		DayID:   dayID,
		HabitID: habitID,
	}
}
