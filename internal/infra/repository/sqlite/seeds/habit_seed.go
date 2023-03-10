package seeds

import (
	"time"

	"github.com/waliqueiroz/habits-api/internal/domain"
	"github.com/waliqueiroz/habits-api/pkg/identity"
)

var firstHabitId = identity.NewULID()
var secondHabitId = identity.NewULID()
var thirdHabitId = identity.NewULID()

func (s Seed) HabitSeed() {
	firstHabitCreationDate, _ := time.Parse(time.RFC3339Nano, "2022-12-31T03:00:00.000Z")
	secondHabitCreationDate, _ := time.Parse(time.RFC3339Nano, "2023-01-03T03:00:00.000Z")
	thirdHabitCreationDate, _ := time.Parse(time.RFC3339Nano, "2023-01-08T03:00:00.000Z")

	habits := []domain.Habit{
		{
			ID:    firstHabitId,
			Title: "Beber 2L de água",
			Weekdays: []domain.HabitWeekday{
				{
					Weekday: 1,
				},
				{
					Weekday: 2,
				},
				{
					Weekday: 3,
				},
			},
			CreatedAt: firstHabitCreationDate,
		},
		{
			ID:    secondHabitId,
			Title: "Exercitar",
			Weekdays: []domain.HabitWeekday{
				{
					Weekday: 3,
				},
				{
					Weekday: 4,
				},
				{
					Weekday: 5,
				},
			},
			CreatedAt: secondHabitCreationDate,
		},
		{
			ID:    thirdHabitId,
			Title: "Dormir 8h",
			Weekdays: []domain.HabitWeekday{
				{
					Weekday: 1,
				},
				{
					Weekday: 2,
				},
				{
					Weekday: 3,
				},
				{
					Weekday: 4,
				},
				{
					Weekday: 5,
				},
			},
			CreatedAt: thirdHabitCreationDate,
		},
	}

	for _, habit := range habits {
		// prepare the statement
		stmt, err := s.db.Prepare(`INSERT INTO habits(id, title, created_at) VALUES (?,?,?)`)
		if err != nil {
			panic(err)
		}
		// execute query
		_, err = stmt.Exec(habit.ID, habit.Title, habit.CreatedAt)
		if err != nil {
			panic(err)
		}

		for _, weekday := range habit.Weekdays {
			// prepare the statement
			stmt, err := s.db.Prepare(`INSERT INTO habit_weekdays(id, habit_id, weekday) VALUES (?,?,?)`)
			if err != nil {
				panic(err)
			}
			// execute query
			_, err = stmt.Exec(identity.NewULID(), habit.ID, weekday.Weekday)
			if err != nil {
				panic(err)
			}
		}
	}
}
