package domain

import "time"

type Habit struct {
	ID        string
	Title     string
	CreatedAt time.Time
}
