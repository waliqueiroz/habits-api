package domain

import "time"

type Day struct {
	ID   string
	Date time.Time
}

type DayHabit struct {
	ID      string
	DayID   string
	HabitID string
}
