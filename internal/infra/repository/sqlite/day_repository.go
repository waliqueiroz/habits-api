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

type DailySummary struct {
	Day
	Completed int `db:"completed"`
	Amount    int `db:"amount"`
}

func mapDailySummaryToDomain(daySummary DailySummary) *domain.DailySummary {
	return &domain.DailySummary{
		Day: domain.Day{
			ID:   daySummary.ID,
			Date: daySummary.Date,
		},
		Completed: daySummary.Completed,
		Amount:    daySummary.Amount,
	}
}

func mapDailySummariesToDomain(daySummaries []DailySummary) []domain.DailySummary {
	list := make([]domain.DailySummary, len(daySummaries))

	for i, daySummary := range daySummaries {
		list[i] = *mapDailySummaryToDomain(daySummary)
	}

	return list

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

func (r *dayRepository) GetSummary(ctx context.Context) ([]domain.DailySummary, error) {
	query := `SELECT
				d.id,
				d.date,
					(
				SELECT
					count (*)
				FROM
					day_habits dh
				WHERE
					dh.day_id = d.id
			) as completed,
				(
				SELECT
					count (*)
				FROM
					habit_weekdays hwd
				JOIN habits h ON
					h.id = hwd.habit_id
				WHERE
					hwd.weekday = cast(strftime('%w', d.date) as int)
						AND h.created_at <= d.date) as amount
			FROM
				days d`

	var summary []DailySummary

	err := r.db.SelectContext(ctx, &summary, query)
	if err != nil {
		return nil, err
	}

	return mapDailySummariesToDomain(summary), nil
}
