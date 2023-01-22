package domain

import "time"

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
