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
	ID        string    `db:"id"`
	Title     string    `db:"title"`
	CreatedAt time.Time `db:"created_at"`
}

func mapHabitToDomain(habit Habit) domain.Habit {
	return domain.Habit{
		ID:        habit.ID,
		Title:     habit.Title,
		CreatedAt: habit.CreatedAt,
	}
}

func mapHabitsToDomain(habits []Habit) []domain.Habit {
	list := make([]domain.Habit, len(habits))

	for i, habit := range habits {
		list[i] = mapHabitToDomain(habit)
	}

	return list
}

type habitRepository struct {
	db *sqlx.DB
}

func NewHabitRepository(db *sqlx.DB) domain.HabitRepository {
	return &habitRepository{
		db: db,
	}
}

func (r *habitRepository) Create(ctx context.Context, habit domain.Habit) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	insertHabitQuery := "INSERT INTO habits(id, title, created_at) VALUES (?,?,?)"

	_, err = tx.ExecContext(ctx, insertHabitQuery, habit.ID, habit.Title, habit.CreatedAt)
	if err != nil {
		return err
	}

	if len(habit.Weekdays) > 0 {
		args := []any{}
		insertWeekDaysAppend := []string{}

		for _, weekday := range habit.Weekdays {
			args = append(args, weekday.ID, weekday.HabitID, weekday.Weekday)
			insertWeekDaysAppend = append(insertWeekDaysAppend, "(?,?,?)")
		}

		insertWeekDaysQuery := fmt.Sprintf("INSERT INTO habit_weekdays(id, habit_id, weekday) VALUES %s", strings.Join(insertWeekDaysAppend, ","))

		_, err = tx.ExecContext(ctx, insertWeekDaysQuery, args...)
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *habitRepository) GetPossibleHabits(ctx context.Context, date time.Time) ([]domain.Habit, error) {
	query := `SELECT * FROM habits h WHERE h.created_at <= ? AND h.id in (SELECT h2.id FROM habits h2 JOIN habit_weekdays hw ON h2.id = hw.habit_id WHERE hw.weekday = ?)`

	habits := []Habit{}

	err := r.db.SelectContext(ctx, &habits, query, date, int(date.Weekday()))
	if err != nil {
		return nil, err
	}

	return mapHabitsToDomain(habits), nil
}

func (r *habitRepository) GetCompletedHabits(ctx context.Context, date time.Time) ([]domain.Habit, error) {
	query := `SELECT
					h.*
				FROM
					habits h
				JOIN day_habits dh on
					h.id = dh.habit_id
				JOIN days d on
					d.id = dh.day_id
				WHERE
					DATE(d.date) = ?`

	habits := []Habit{}

	err := r.db.SelectContext(ctx, &habits, query, date.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}

	return mapHabitsToDomain(habits), nil
}
