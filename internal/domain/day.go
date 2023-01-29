package domain

import (
	"context"
	"time"

	"github.com/waliqueiroz/habits-api/pkg/identity"
)

type DayRepository interface {
	Create(ctx context.Context, day Day) error
	FindByDate(ctx context.Context, date time.Time) (*Day, error)
	GetSummary(ctx context.Context) ([]DailySummary, error)
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

type DailySummary struct {
	Day
	Completed int
	Amount    int
}
