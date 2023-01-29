package application

import (
	"context"
	"time"

	"github.com/waliqueiroz/habits-api/internal/domain"
)

type HabitService interface {
	Create(ctx context.Context, habit domain.Habit) error
	GetDayResume(ctx context.Context, date time.Time) (*domain.DayResume, error)
	ToggleHabit(ctx context.Context, habitID string) error
}

type habitService struct {
	habitRepository    domain.HabitRepository
	dayRepository      domain.DayRepository
	dayHabitRepository domain.DayHabitRepository
}

func NewHabitService(habitRepository domain.HabitRepository, dayRepository domain.DayRepository, dayHabitRepository domain.DayHabitRepository) HabitService {
	return &habitService{
		habitRepository:    habitRepository,
		dayRepository:      dayRepository,
		dayHabitRepository: dayHabitRepository,
	}
}

func (s *habitService) Create(ctx context.Context, habit domain.Habit) error {
	// TODO - executar validação

	err := s.habitRepository.Create(ctx, habit)
	if err != nil {
		return err
	}

	return nil
}

func (s *habitService) GetDayResume(ctx context.Context, date time.Time) (*domain.DayResume, error) {
	possibleHabits, err := s.habitRepository.GetPossibleHabits(ctx, date)
	if err != nil {
		return nil, err
	}

	completedHabits, err := s.habitRepository.GetCompletedHabits(ctx, date)
	if err != nil {
		return nil, err
	}

	return &domain.DayResume{
		PossibleHabits:  possibleHabits,
		CompletedHabits: completedHabits,
	}, nil
}

func (s *habitService) ToggleHabit(ctx context.Context, habitID string) error {
	today := domain.TruncateToDay(time.Now())

	day, err := s.dayRepository.FindByDate(ctx, today)
	if err != nil {
		switch err {
		case domain.ErrResourceNotFound:
			err := s.dayRepository.Create(ctx, domain.NewDay(today))
			if err != nil {
				return err
			}
		default:
			return err
		}
	}

	existsDayHabit, err := s.dayHabitRepository.ExistsByDayIDAndHabitID(ctx, day.ID, habitID)
	if err != nil {
		return err
	}

	if existsDayHabit {
		err := s.dayHabitRepository.DeleteByDayIDAndHabitID(ctx, day.ID, habitID)
		if err != nil {
			return err
		}
	} else {
		err := s.dayHabitRepository.Create(ctx, domain.NewDayHabit(day.ID, habitID))
		if err != nil {
			return err
		}
	}

	return nil
}
