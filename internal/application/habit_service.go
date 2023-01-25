package application

import (
	"context"
	"time"

	"github.com/waliqueiroz/habits-api/internal/domain"
)

type HabitService interface {
	Create(ctx context.Context, habit domain.Habit) error
	GetDayResume(ctx context.Context, date time.Time) (*domain.DayResume, error)
}

type habitService struct {
	habitRepository domain.HabitRepository
}

func NewHabitService(habitRepository domain.HabitRepository) HabitService {
	return &habitService{
		habitRepository: habitRepository,
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
