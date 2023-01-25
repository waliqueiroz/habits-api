package seeds

import (
	"time"

	"github.com/waliqueiroz/habits-api/internal/domain"
	"github.com/waliqueiroz/habits-api/pkg/identity"
)

func (s Seed) DaySeed() {
	dayOne, _ := time.Parse(time.RFC3339Nano, "2023-01-02T03:00:00.000Z")
	dayTwo, _ := time.Parse(time.RFC3339Nano, "2023-01-06T03:00:00.000Z")
	dayThree, _ := time.Parse(time.RFC3339Nano, "2023-01-04T03:00:00.000Z")

	days := []domain.Day{
		{
			ID:   identity.NewULID(),
			Date: dayOne,
			Habits: []domain.Habit{
				{
					ID: firstHabitId,
				},
			},
		},
		{
			ID:   identity.NewULID(),
			Date: dayTwo,
			Habits: []domain.Habit{
				{
					ID: firstHabitId,
				},
			},
		},
		{
			ID:   identity.NewULID(),
			Date: dayThree,
			Habits: []domain.Habit{
				{
					ID: firstHabitId,
				},

				{
					ID: secondHabitId,
				},
			},
		},
	}

	for _, day := range days {
		// prepare the statement
		stmt, err := s.db.Prepare(`INSERT INTO days(id, date) VALUES (?,?)`)
		if err != nil {
			panic(err)
		}
		// execute query
		_, err = stmt.Exec(day.ID, day.Date)
		if err != nil {
			panic(err)
		}

		for _, habit := range day.Habits {
			// prepare the statement
			stmt, err := s.db.Prepare(`INSERT INTO day_habits(id, day_id, habit_id) VALUES (?,?,?)`)
			if err != nil {
				panic(err)
			}
			// execute query
			_, err = stmt.Exec(identity.NewULID(), day.ID, habit.ID)
			if err != nil {
				panic(err)
			}
		}
	}
}
