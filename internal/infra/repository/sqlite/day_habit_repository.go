package sqlite

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/waliqueiroz/habits-api/internal/domain"
)

type dayHabitRepository struct {
	db *sqlx.DB
}

func NewDayHabitRepository(db *sqlx.DB) domain.DayHabitRepository {
	return &dayHabitRepository{
		db: db,
	}
}

func (r *dayHabitRepository) Create(ctx context.Context, dayHabit domain.DayHabit) error {
	query := "INSERT INTO day_habits(id, day_id, habit_id) VALUES (?,?,?)"

	_, err := r.db.ExecContext(ctx, query, dayHabit.ID, dayHabit.DayID, dayHabit.HabitID)
	if err != nil {
		return err
	}

	return nil
}

func (r *dayHabitRepository) ExistsByDayIDAndHabitID(ctx context.Context, dayID string, habitID string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 from day_habits dh WHERE dh.day_id = ? AND dh.habit_id = ?)`

	var exists bool

	err := r.db.GetContext(ctx, &exists, query, dayID, habitID)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *dayHabitRepository) DeleteByDayIDAndHabitID(ctx context.Context, dayID string, habitID string) error {
	query := `DELETE from day_habits WHERE day_id = ? AND habit_id = ?`

	_, err := r.db.ExecContext(ctx, query, dayID, habitID)
	if err != nil {
		return err
	}

	return nil
}
