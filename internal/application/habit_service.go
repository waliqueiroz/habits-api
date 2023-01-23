package application

import (
	"context"

	"github.com/waliqueiroz/habits-api/internal/domain"
)

type HabitService interface {
	Create(ctx context.Context, habit domain.Habit) (*domain.Habit, error)
}

type habitService struct {
	habitRepository domain.HabitRepository
}

func NewHabitService(habitRepository domain.HabitRepository) HabitService {
	return &habitService{
		habitRepository: habitRepository,
	}
}

func (s *habitService) Create(ctx context.Context, habit domain.Habit) (*domain.Habit, error) {
	// TODO - executar validação

	newHabit, err := s.habitRepository.Create(ctx, habit)
	if err != nil {
		return nil, err
	}

	return newHabit, nil
}
