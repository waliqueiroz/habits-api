package sqlite

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/waliqueiroz/habits-api/internal/domain"
)

type Habit struct {
	ID        string         `db:"id"`
	Title     string         `db:"title"`
	Weekdays  []HabitWeekday `db:"weekdays"`
	CreatedAt time.Time      `db:"created_at"`
}

type HabitWeekday struct {
	ID      string `db:"id"`
	HabitID string `db:"habit_id"`
	Weekday int    `db:"weekday"`
}

type habitRepository struct {
	db *sqlx.DB
}

func NewHabitRepository(db *sqlx.DB) domain.HabitRepository {
	return &habitRepository{
		db: db,
	}
}

func (r *habitRepository) Create(ctx context.Context, habit domain.Habit) (*domain.Habit, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	insertHabit := "INSERT INTO habits(id, title, created_at) VALUES (?,?,?)"

	_, err = tx.ExecContext(ctx, insertHabit, habit.ID, habit.Title, habit.CreatedAt)
	if err != nil {
		return nil, err
	}

	if len(habit.Weekdays) > 0 {
		args := []any{}
		insertWeekDaysAppend := []string{}

		for _, weekday := range habit.Weekdays {
			args = append(args, weekday.ID, weekday.HabitID, weekday.Weekday)
			insertWeekDaysAppend = append(insertWeekDaysAppend, "(?,?,?)")
		}

		insertWeekDays := fmt.Sprintf("INSERT INTO habit_weekdays(id, habit_id, weekday) VALUES %s", strings.Join(insertWeekDaysAppend, ","))

		_, err = tx.ExecContext(ctx, insertWeekDays, args...)
		if err != nil {
			return nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &habit, nil
}
