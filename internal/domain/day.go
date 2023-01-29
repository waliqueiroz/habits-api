package domain

import (
	"context"
	"time"

	"github.com/waliqueiroz/habits-api/pkg/identity"
)

type DayRepository interface {
	Create(ctx context.Context, day Day) error
	FindByDate(ctx context.Context, date time.Time) (*Day, error)
}

type Day struct {
	ID     string
	Date   time.Time
	Habits []Habit
}

func NewDay(date time.Time) Day {
	return Day{
		ID:   identity.NewULID(),
		Date: date,
	}
}
