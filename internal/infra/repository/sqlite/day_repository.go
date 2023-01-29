package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/waliqueiroz/habits-api/internal/domain"
)

type Day struct {
	ID   string    `db:"id"`
	Date time.Time `db:"date"`
}

func mapDayToDomain(day Day) *domain.Day {
	return &domain.Day{
		ID:   day.ID,
		Date: day.Date,
	}
}

type dayRepository struct {
	db *sqlx.DB
}

func NewDayRepository(db *sqlx.DB) domain.DayRepository {
	return &dayRepository{
		db: db,
	}
}

func (r *dayRepository) Create(ctx context.Context, day domain.Day) error {
	query := "INSERT INTO days(id, date) VALUES (?,?)"

	_, err := r.db.ExecContext(ctx, query, day.ID, day.Date)
	if err != nil {
		return err
	}

	return nil
}

func (r *dayRepository) FindByDate(ctx context.Context, date time.Time) (*domain.Day, error) {
	query := `SELECT * FROM days d WHERE d.date = ?`

	var day Day

	err := r.db.GetContext(ctx, &day, query, date)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrResourceNotFound
		}
		return nil, err
	}

	return mapDayToDomain(day), nil
}
